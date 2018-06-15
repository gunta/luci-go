// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repo

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/gae/service/datastore"

	"go.chromium.org/luci/appengine/tq"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/iotools"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/server/auth"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
	"go.chromium.org/luci/cipd/appengine/impl/cas"
	"go.chromium.org/luci/cipd/appengine/impl/metadata"
	"go.chromium.org/luci/cipd/appengine/impl/model"
	"go.chromium.org/luci/cipd/appengine/impl/repo/processing"
	"go.chromium.org/luci/cipd/appengine/impl/repo/tasks"
	"go.chromium.org/luci/cipd/common"
)

// Public returns publicly exposed implementation of cipd.Repository service.
//
// It checks ACLs.
func Public(internalCAS cas.StorageServer, d *tq.Dispatcher) api.RepositoryServer {
	impl := &repoImpl{
		tq:   d,
		meta: metadata.GetStorage(),
		cas:  internalCAS,
	}
	impl.registerTasks()
	impl.registerProcessor(&processing.ClientExtractor{CAS: internalCAS})
	return impl
}

// repoImpl implements api.RepositoryServer.
type repoImpl struct {
	tq *tq.Dispatcher

	meta metadata.Storage  // storage for package prefix metadata
	cas  cas.StorageServer // non-ACLed storage for instance package files

	procs    []processing.Processor          // in order of registerProcessor calls
	procsMap map[string]processing.Processor // ID => processing.Processor
}

// registerTasks adds tasks to the tq Dispatcher.
func (impl *repoImpl) registerTasks() {
	// See queue.yaml for "run-processors" task queue definition.
	impl.tq.RegisterTask(&tasks.RunProcessors{}, func(c context.Context, m proto.Message) error {
		return impl.runProcessorsTask(c, m.(*tasks.RunProcessors))
	}, "run-processors", nil)
}

// registerProcessor adds a new processor.
func (impl *repoImpl) registerProcessor(p processing.Processor) {
	if impl.procsMap == nil {
		impl.procsMap = map[string]processing.Processor{}
	}

	id := p.ID()
	if impl.procsMap[id] != nil {
		panic(fmt.Sprintf("processor %q has already been registered", id))
	}

	impl.procs = append(impl.procs, p)
	impl.procsMap[id] = p
}

// packageReader opens a package instance for reading.
func (impl *repoImpl) packageReader(c context.Context, ref *api.ObjectRef) (*processing.PackageReader, error) {
	// Get slow Google Storage based ReaderAt.
	rawReader, err := impl.cas.GetReader(c, ref)
	switch code := grpc.Code(err); {
	case code == codes.NotFound:
		return nil, errors.Annotate(err, "package instance is not in the storage").Err()
	case code != codes.OK:
		return nil, errors.Annotate(err, "failed to open the object for reading").Tag(transient.Tag).Err()
	}

	// Read in 512 Kb chunks, keep 2 of them buffered.
	pkg, err := processing.NewPackageReader(
		iotools.NewBufferingReaderAt(rawReader, 512*1024, 2),
		rawReader.Size())
	if err != nil {
		return nil, errors.Annotate(err, "error when opening the package").Err()
	}
	return pkg, nil
}

////////////////////////////////////////////////////////////////////////////////
// Prefix metadata RPC methods + related helpers including ACL checks.

// GetPrefixMetadata implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) GetPrefixMetadata(c context.Context, r *api.PrefixRequest) (resp *api.PrefixMetadata, err error) {
	// It is fine to implement this in terms of GetInheritedPrefixMetadata, since
	// we need to fetch all inherited metadata anyway to check ACLs.
	inherited, err := impl.GetInheritedPrefixMetadata(c, r)
	if err != nil {
		return nil, err
	}
	// Have the metadata for the requested prefix? It should be the last if so.
	if m := inherited.PerPrefixMetadata; len(m) != 0 && m[len(m)-1].Prefix == r.Prefix {
		return m[len(m)-1], nil
	}
	// Note that GetInheritedPrefixMetadata checked that the caller has permission
	// to view the requested prefix (via some parent prefix ACL), so sincerely
	// reply with NotFound.
	return nil, noMetadataErr(r.Prefix)
}

// GetInheritedPrefixMetadata implements the corresponding RPC method, see the
// proto doc.
//
// Note: it normalizes Prefix field inside the request.
func (impl *repoImpl) GetInheritedPrefixMetadata(c context.Context, r *api.PrefixRequest) (resp *api.InheritedPrefixMetadata, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	r.Prefix, err = common.ValidatePackagePrefix(r.Prefix)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'prefix' - %s", err)
	}

	metas, err := impl.checkRole(c, r.Prefix, api.Role_OWNER)
	if err != nil {
		return nil, err
	}
	return &api.InheritedPrefixMetadata{PerPrefixMetadata: metas}, nil
}

// UpdatePrefixMetadata implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) UpdatePrefixMetadata(c context.Context, r *api.PrefixMetadata) (resp *api.PrefixMetadata, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Fill in server-assigned fields.
	r.UpdateTime = google.NewTimestamp(clock.Now(c))
	r.UpdateUser = string(auth.CurrentIdentity(c))

	// Normalize and validate format of the PrefixMetadata.
	if err := common.NormalizePrefixMetadata(r); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad prefix metadata - %s", err)
	}

	// The root metadata is not modifiable through API.
	if r.Prefix == "" {
		return nil, status.Errorf(codes.InvalidArgument, "the root metadata is not modifiable")
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Prefix, api.Role_OWNER); err != nil {
		return nil, err
	}

	// Transactionally check the fingerprint and update the metadata. impl.meta
	// will recalculate the new fingerprint. Note there's a small chance the
	// caller no longer has OWNER role to modify the metadata inside the
	// transaction. We ignore it. It happens when caller's permissions are revoked
	// by someone else exactly during UpdatePrefixMetadata call.
	return impl.meta.UpdateMetadata(c, r.Prefix, func(cur *api.PrefixMetadata) error {
		if cur.Fingerprint != r.Fingerprint {
			switch {
			case cur.Fingerprint == "":
				// The metadata was deleted while the caller was messing with it.
				return noMetadataErr(r.Prefix)
			case r.Fingerprint == "":
				// Caller tries to make a new one, but we already have it.
				return status.Errorf(
					codes.AlreadyExists, "metadata for prefix %q already exists and has fingerprint %q, "+
						"use combination of GetPrefixMetadata and UpdatePrefixMetadata to "+
						"update it", r.Prefix, cur.Fingerprint)
			default:
				// The fingerprint has changed while the caller was messing with
				// the metadata.
				return status.Errorf(
					codes.FailedPrecondition, "metadata for prefix %q was updated concurrently "+
						"(the fingerprint in the request %q doesn't match the current fingerprint %q), "+
						"fetch new metadata with GetPrefixMetadata and reapply your "+
						"changes", r.Prefix, r.Fingerprint, cur.Fingerprint)
			}
		}
		*cur = *r
		return nil
	})
}

// GetRolesInPrefix implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) GetRolesInPrefix(c context.Context, r *api.PrefixRequest) (resp *api.RolesInPrefixResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	r.Prefix, err = common.ValidatePackagePrefix(r.Prefix)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'prefix' - %s", err)
	}

	metas, err := impl.meta.GetMetadata(c, r.Prefix)
	if err != nil {
		return nil, err
	}

	roles, err := rolesInPrefix(c, metas)
	if err != nil {
		return nil, err
	}

	resp = &api.RolesInPrefixResponse{
		Roles: make([]*api.RolesInPrefixResponse_RoleInPrefix, len(roles)),
	}
	for i, r := range roles {
		resp.Roles[i] = &api.RolesInPrefixResponse_RoleInPrefix{Role: r}
	}
	return resp, nil
}

// checkRole checks where the caller has the given role in the given prefix or
// any of its parent prefixes.
//
// Understands role inheritance. See acl.go for more details.
//
// Returns grpc PermissionDenied error if the caller doesn't have the requested
// role. The error message depends on whether caller has READER role or not
// (readers see more details).
//
// Fetches and returns metadata of the prefix and all parent prefixes as a side
// effect.
func (impl *repoImpl) checkRole(c context.Context, prefix string, role api.Role) ([]*api.PrefixMetadata, error) {
	metas, err := impl.meta.GetMetadata(c, prefix)
	if err != nil {
		return nil, err
	}

	switch yes, err := hasRole(c, metas, role); {
	case err != nil:
		return nil, err
	case yes:
		return metas, nil
	case role == api.Role_READER: // was checking for a reader, and caller is not
		return nil, noAccessErr(prefix)
	}

	// We end up here if role is something other than READER, and the caller
	// doesn't have it. Maybe caller IS a reader, then we can give more concrete
	// error message.
	switch yes, err := hasRole(c, metas, api.Role_READER); {
	case err != nil:
		return nil, err
	case yes:
		return nil, status.Errorf(codes.PermissionDenied, "caller has no required %s role in prefix %q", role, prefix)
	default:
		return nil, noAccessErr(prefix)
	}
}

// noAccessErr produces a grpc error saying that the given prefix doesn't
// exist or the caller has no access to it. This is generic error message that
// should not give away prefix presence to non-readers.
func noAccessErr(prefix string) error {
	return status.Errorf(codes.PermissionDenied, "prefix %q doesn't exist or the caller is not allowed to see it", prefix)
}

// noMetadataErr produces a grpc error saying that the given prefix doesn't have
// metadata attached.
func noMetadataErr(prefix string) error {
	return status.Errorf(codes.NotFound, "prefix %q has no metadata", prefix)
}

////////////////////////////////////////////////////////////////////////////////
// Prefix listing.

// ListPrefix implement the corresponding RPC method, see the proto doc.
func (impl *repoImpl) ListPrefix(c context.Context, r *api.ListPrefixRequest) (resp *api.ListPrefixResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	r.Prefix, err = common.ValidatePackagePrefix(r.Prefix)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'prefix' - %s", err)
	}

	// Discover prefixes the caller is allowed to see. Note that checking only
	// r.Prefix ACL is not sufficient, since the caller may not see it, but still
	// see some deeper prefix, thus we need to enumerate the metadata subtree.
	var visibleRoots []string // sorted list of prefixes visible to the caller
	err = impl.meta.VisitMetadata(c, r.Prefix, func(pfx string, md []*api.PrefixMetadata) (cont bool, err error) {
		switch visible, err := hasRole(c, md, api.Role_READER); {
		case err != nil:
			return false, err
		case visible:
			// Found a visible root. Everything under 'pfx' is visible to the caller,
			// no sense in recursing deeper into this subtree.
			visibleRoots = append(visibleRoots, pfx)
			return false, nil
		default:
			// Continue exploring this subtree until we find something visible there,
			// if anything.
			return true, nil
		}
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to enumerate the metadata").Err()
	}

	// One of 'visibleRoots' => its full sorted recursive listing.
	perVisibleRoot := make(map[string][]string, len(visibleRoots))

	// ListPackages used below lists only packages that are under the prefix. It
	// is possible there are packages that match some of visibleRoots directly
	// (have exact same name as some visible root). We check it here. Note that
	// this applies only to subprefixes of r.Prefix: we don't want to return a
	// package named r.Prefix in the result (if any). So if visibleRoots is
	// [r.Prefix] itself, we skip this check. Note that if r.Prefix is in
	// visibleRoots, then it is the only item there, by construction.
	if len(visibleRoots) != 1 || visibleRoots[0] != r.Prefix {
		rootPkgs, err := model.CheckPackages(c, visibleRoots, r.IncludeHidden)
		if err != nil {
			return nil, errors.Annotate(err, "failed to check presence of packages").Err()
		}
		for _, name := range rootPkgs {
			perVisibleRoot[name] = []string{name}
		}
	}

	// Fetch the listing of each visible root in parallel. Note that there's no
	// intersection between them due to the order of VisitMetadata enumeration.
	err = parallel.WorkPool(8, func(tasks chan<- func() error) {
		mu := sync.Mutex{}
		for _, pfx := range visibleRoots {
			pfx := pfx
			tasks <- func() error {
				// TODO(vadimsh): This is inefficient for non-recursive listings.
				// Unfortunately we have to do the recursive listing to discover
				// subprefixes. E.g. when listing 'a', we need to return 'a/b' as a
				// subprefix if there's a package 'a/b/c', so to discover 'a/b/c' we
				// need full recursive listing of 'a'.
				//
				// One possible improvement is:
				//  1. Introduce a new parallel entity group just for browsing:
				//     Entry {
				//       Name: <full path to the package or the prefix>,
				//       Children: [<names of all child prefixes and packages>],
				//       HiddenChildren: [<names of hidden prefixes and packages>],
				//       Hidden: <true|false>,  // also true if all children are hidden
				//     }
				//  2. 'Entry' basically represents a browsable item (either a package
				//     or a "directory" aka prefix).
				//  3. Update this entity group when registering, deleting, hiding or
				//     showing packages. Will require O(<number of path components>)
				//     entity writes in a single transaction, which is acceptable.
				//  4. In non-recursive listing just look at Children of corresponding
				//     entity.
				listing, err := model.ListPackages(c, pfx, r.IncludeHidden)
				if err == nil {
					mu.Lock()
					perVisibleRoot[pfx] = append(perVisibleRoot[pfx], listing...)
					mu.Unlock()
				}
				return err
			}
		}
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to list a prefix").Err()
	}

	// searchRoot is lexicographical prefix of packages we are concerned about.
	searchRoot := r.Prefix
	if searchRoot != "" {
		searchRoot += "/"
	}

	// Visit all discovered packages (in sorted order).
	resp = &api.ListPrefixResponse{}
	dirs := stringset.New(0)
	for _, pfx := range visibleRoots {
		for _, pkg := range perVisibleRoot[pfx] {
			// By construction, ALL packages here must be within the search root.
			if !strings.HasPrefix(pkg, searchRoot) {
				panic(fmt.Sprintf(
					"unexpected package %q in results when listing %q, visible roots are %q",
					pkg, searchRoot, visibleRoots))
			}

			// E.g. "a/b/c", relative to searchRoot.
			rel := strings.TrimPrefix(pkg, searchRoot)

			if r.Recursive {
				// If listing recursively, add everything to the result.
				resp.Packages = append(resp.Packages, pkg)
				// For rel "a/b/c", add [".../a", ".../a/b"] to the directories set.
				if chunks := strings.Split(rel, "/"); len(chunks) > 1 {
					for i := 1; i < len(chunks); i++ {
						dirs.Add(searchRoot + strings.Join(chunks[:i], "/"))
					}
				}
			} else {
				// Otherwise add only packages that directly reside under searchRoot,
				// and pick only first path component of directories, since we don't
				// care about what's deeper.
				if idx := strings.IndexRune(rel, '/'); idx != -1 {
					// 'rel' has / => it is a directory, add its first path component.
					dirs.Add(searchRoot + rel[:idx])
				} else {
					// 'rel' has no / inside => it's a package directly under searchRoot.
					resp.Packages = append(resp.Packages, pkg)
				}
			}
		}
	}

	// Note: resp.Packages are already sorted by construction.
	resp.Prefixes = dirs.ToSlice()
	sort.Strings(resp.Prefixes)

	return resp, nil
}

////////////////////////////////////////////////////////////////////////////////
// Hide/unhide package.

// HidePackage implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) HidePackage(c context.Context, r *api.PackageRequest) (*empty.Empty, error) {
	return impl.setPackageHidden(c, r, model.Hidden)
}

// UnhidePackage implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) UnhidePackage(c context.Context, r *api.PackageRequest) (*empty.Empty, error) {
	return impl.setPackageHidden(c, r, model.Visible)
}

// setPackageHidden is common implementation of HidePackage and UnhidePackage.
func (impl *repoImpl) setPackageHidden(c context.Context, r *api.PackageRequest, hidden bool) (resp *empty.Empty, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if _, err := impl.checkRole(c, r.Package, api.Role_OWNER); err != nil {
		return nil, err
	}

	switch err := model.SetPackageHidden(c, r.Package, hidden); {
	case err == datastore.ErrNoSuchEntity:
		return nil, status.Errorf(codes.NotFound, "no such package")
	case err != nil:
		return nil, errors.Annotate(err, "failed to update the package").Err()
	}
	return &empty.Empty{}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Package instance registration and post-registration processing.

// RegisterInstance implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) RegisterInstance(c context.Context, r *api.Instance) (resp *api.RegisterInstanceResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request format.
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if err := cas.ValidateObjectRef(r.Instance); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'instance' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_WRITER); err != nil {
		return nil, err
	}

	// Is such instance already registered?
	instance := (&model.Instance{}).FromProto(c, r)
	switch err := datastore.Get(c, instance); {
	case err == nil:
		return &api.RegisterInstanceResponse{
			Status:   api.RegistrationStatus_ALREADY_REGISTERED,
			Instance: instance.Proto(),
		}, nil
	case err != datastore.ErrNoSuchEntity:
		return nil, errors.Annotate(err, "failed to fetch the instance entity").Err()
	}

	// Attempt to start a new upload session. This will fail with ALREADY_EXISTS
	// if such object is already in the storage. This is expected (it means the
	// client has uploaded the object already and we should just register the
	// instance right away).
	uploadOp, err := impl.cas.BeginUpload(c, &api.BeginUploadRequest{
		Object: r.Instance,
	})
	switch code := grpc.Code(err); {
	case code == codes.AlreadyExists:
		break // the object is already there
	case code == codes.OK:
		// The object is not in the storage and we have just started the upload. Let
		// the client finish it.
		return &api.RegisterInstanceResponse{
			Status:   api.RegistrationStatus_NOT_UPLOADED,
			UploadOp: uploadOp,
		}, nil
	default:
		return nil, errors.Annotate(err, "failed to initiate an upload op (code %s)", code).Err()
	}

	// The instance is already in the CAS storage. Register it in the repository.
	instance = (&model.Instance{
		RegisteredBy: string(auth.CurrentIdentity(c)),
		RegisteredTs: clock.Now(c).UTC(),
	}).FromProto(c, r)
	registered, instance, err := model.RegisterInstance(c, instance, impl.onInstanceRegistration)
	if err != nil {
		return nil, errors.Annotate(err, "failed to register the instance").Err()
	}

	resp = &api.RegisterInstanceResponse{Instance: instance.Proto()}
	if registered {
		resp.Status = api.RegistrationStatus_REGISTERED
	} else {
		resp.Status = api.RegistrationStatus_ALREADY_REGISTERED
	}
	return
}

// onInstanceRegistration is called in a txn when registering an instance.
func (impl *repoImpl) onInstanceRegistration(c context.Context, inst *model.Instance) error {
	// Collect IDs of applicable processors.
	var procs []string
	for _, p := range impl.procs {
		if p.Applicable(inst) {
			procs = append(procs, p.ID())
		}
	}
	if len(procs) == 0 {
		return nil
	}

	// Mark the instance as being processed now.
	inst.ProcessorsPending = procs

	// Launch the TQ task that does the processing (see runProcessorsTask below).
	return impl.tq.AddTask(c, &tq.Task{
		Payload: &tasks.RunProcessors{Instance: inst.Proto()},
		Title:   inst.InstanceID,
	})
}

// runProcessorsTask executes a post-upload processing step.
//
// Returning a transient error here causes the task queue service to retry the
// task.
func (impl *repoImpl) runProcessorsTask(c context.Context, t *tasks.RunProcessors) error {
	// Fetch the instance to see what processors are still pending.
	inst := (&model.Instance{}).FromProto(c, t.Instance)
	switch err := datastore.Get(c, inst); {
	case err == datastore.ErrNoSuchEntity:
		return fmt.Errorf("instance %q is unexpectedly gone from the datastore", inst.InstanceID)
	case err != nil:
		return transient.Tag.Apply(err)
	}

	results := map[string]processing.Result{}

	// Grab processors we haven't ran yet.
	var run []processing.Processor
	for _, id := range inst.ProcessorsPending {
		if proc := impl.procsMap[id]; proc != nil {
			run = append(run, proc)
		} else {
			logging.Errorf(c, "Skipping unknown processor %q", id)
			results[id] = processing.Result{Err: fmt.Errorf("unknown processor %q", id)}
		}
	}

	// Exit early if there's nothing to run.
	if len(run) == 0 {
		return impl.updateProcessors(c, t.Instance, results)
	}

	// Open the package for reading.
	pkg, err := impl.packageReader(c, t.Instance.Instance)
	switch {
	case transient.Tag.In(err):
		return err // retry the whole thing
	case err != nil:
		// The package is fatally broken, give up.
		logging.WithError(err).Errorf(c, "The package can't be opened, failing all processors")
		for _, proc := range run {
			results[proc.ID()] = processing.Result{Err: err}
		}
		return impl.updateProcessors(c, t.Instance, results)
	}

	// Run the processors sequentially, since PackageReader is not very friendly
	// to concurrent access.
	var transientErrs errors.MultiError
	for _, proc := range run {
		logging.Infof(c, "Running processor %q", proc.ID())
		res, err := proc.Run(c, inst, pkg)
		if err != nil {
			logging.WithError(err).Errorf(c, "Processor %q failed transiently", proc.ID())
			transientErrs = append(transientErrs, err)
		} else {
			if res.Err != nil {
				logging.WithError(res.Err).Errorf(c, "Processor %q failed fatally", proc.ID())
			}
			results[proc.ID()] = res
		}
	}

	// Store what we've got, even if some processor may have failed to run.
	updErr := impl.updateProcessors(c, t.Instance, results)

	// Prefer errors from processors over 'updErr' if both happen. Processor
	// errors are more interesting.
	switch {
	case len(transientErrs) != 0:
		return transient.Tag.Apply(transientErrs)
	case updErr != nil:
		return updErr
	}
	return nil
}

// updateProcessors transactionally creates ProcessingResult entities and
// updates Instance.Processors* fields.
func (impl *repoImpl) updateProcessors(c context.Context, inst *api.Instance, results map[string]processing.Result) error {
	if len(results) == 0 {
		return nil
	}

	instEnt := (&model.Instance{}).FromProto(c, inst)
	instKey := datastore.KeyForObj(c, instEnt)

	now := clock.Now(c).UTC()

	// Create ProcessingResult outside the transaction, since this involves slow
	// zlib compression in WriteResult.
	procResults := make(map[string]*model.ProcessingResult, len(results))
	for procID, res := range results {
		procRes := &model.ProcessingResult{
			ProcID:    procID,
			Instance:  instKey,
			CreatedTs: now,
			Success:   res.Err == nil,
		}
		procResults[procID] = procRes

		// If the result is not serializable, store the serialization error instead.
		err := res.Err
		if err == nil {
			if err = procRes.WriteResult(res.Result); err != nil {
				err = errors.Annotate(err, "failed to write the processing result").Err()
			}
		}
		if err != nil {
			procRes.Success = false
			procRes.Error = err.Error()
		}
	}

	// Mutate Instance entity, storing results that haven't been stored yet.
	return model.Txn(c, "updateProcessors", func(c context.Context) error {
		switch err := datastore.Get(c, instEnt); {
		case err == datastore.ErrNoSuchEntity:
			return fmt.Errorf("the entity is unexpectedly gone")
		case err != nil:
			return errors.Annotate(err, "failed to fetch the entity").Tag(transient.Tag).Err()
		}

		var toPut []interface{}

		// Go over what's is still pending, and move it to either Success or Failure
		// group if it is done.
		stillPending := instEnt.ProcessorsPending[:0]
		for _, procID := range instEnt.ProcessorsPending {
			res, done := procResults[procID]
			if !done {
				stillPending = append(stillPending, procID)
				continue
			}
			toPut = append(toPut, res)
			if res.Success {
				instEnt.ProcessorsSuccess = append(instEnt.ProcessorsSuccess, procID)
			} else {
				instEnt.ProcessorsFailure = append(instEnt.ProcessorsFailure, procID)
			}
		}
		instEnt.ProcessorsPending = stillPending

		// Store all the changes (if any).
		if len(toPut) == 0 {
			return nil
		}
		return transient.Tag.Apply(datastore.Put(c, toPut, instEnt))
	})
}

////////////////////////////////////////////////////////////////////////////////
// Instance listing.

// ListInstances implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) ListInstances(c context.Context, r *api.ListInstancesRequest) (resp *api.ListInstancesResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request, decode the cursor.
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}

	switch {
	case r.PageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "bad 'page_size' %d - it should be non-negative", r.PageSize)
	case r.PageSize == 0:
		r.PageSize = 100
	}

	var cursor datastore.Cursor
	if r.PageToken != "" {
		if cursor, err = datastore.DecodeCursor(c, r.PageToken); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "bad 'page_token' - %s", err)
		}
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_READER); err != nil {
		return nil, err
	}

	// Check that the package is registered.
	if err := model.CheckPackageExists(c, r.Package); err != nil {
		return nil, err
	}

	// Do the actual listing.
	inst, cursor, err := model.ListInstances(c, r.Package, r.PageSize, cursor)
	if err != nil {
		return nil, errors.Annotate(err, "failed to list instances").Err()
	}

	// Convert results to proto.
	resp = &api.ListInstancesResponse{
		Instances: make([]*api.Instance, len(inst)),
	}
	for i, ent := range inst {
		resp.Instances[i] = ent.Proto()
	}
	if cursor != nil {
		resp.NextPageToken = cursor.String()
	}
	return resp, nil
}

////////////////////////////////////////////////////////////////////////////////
// Refs support.

// CreateRef implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) CreateRef(c context.Context, r *api.Ref) (resp *empty.Empty, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := common.ValidatePackageRef(r.Name); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'name' - %s", err)
	}
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if err := cas.ValidateObjectRef(r.Instance); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'instance' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_WRITER); err != nil {
		return nil, err
	}

	// Actually create or move the ref. This will also transactionally check the
	// instance exists and it has passed the processing successfully.
	inst := &model.Instance{
		InstanceID: model.ObjectRefToInstanceID(r.Instance),
		Package:    model.PackageKey(c, r.Package),
	}
	if err := model.SetRef(c, r.Name, inst, auth.CurrentIdentity(c)); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DeleteRef implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) DeleteRef(c context.Context, r *api.DeleteRefRequest) (resp *empty.Empty, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := common.ValidatePackageRef(r.Name); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'name' - %s", err)
	}
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_WRITER); err != nil {
		return nil, err
	}

	// Verify the package actually exists, per DeleteRef contract.
	if err := model.CheckPackageExists(c, r.Package); err != nil {
		return nil, err
	}

	// Actually delete the ref.
	if err := model.DeleteRef(c, r.Package, r.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ListRefs implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) ListRefs(c context.Context, r *api.ListRefsRequest) (resp *api.ListRefsResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_READER); err != nil {
		return nil, err
	}

	// Verify the package actually exists, per ListPackageRefs contract.
	if err := model.CheckPackageExists(c, r.Package); err != nil {
		return nil, err
	}

	// Actually list refs.
	refs, err := model.ListPackageRefs(c, r.Package)
	if err != nil {
		return nil, errors.Annotate(err, "failed to list refs").Err()
	}
	resp = &api.ListRefsResponse{Refs: make([]*api.Ref, len(refs))}
	for i, ref := range refs {
		resp.Refs[i] = ref.Proto()
	}
	return resp, nil
}

////////////////////////////////////////////////////////////////////////////////
// Tags support.

func validateMultiTagReq(pkg string, inst *api.ObjectRef, tags []*api.Tag) error {
	if err := common.ValidatePackageName(pkg); err != nil {
		return status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if err := cas.ValidateObjectRef(inst); err != nil {
		return status.Errorf(codes.InvalidArgument, "bad 'instance' - %s", err)
	}
	if len(tags) == 0 {
		return status.Errorf(codes.InvalidArgument, "bad 'tags' - cannot be empty")
	}
	for _, t := range tags {
		kv := common.JoinInstanceTag(t)
		if err := common.ValidateInstanceTag(kv); err != nil {
			return status.Errorf(codes.InvalidArgument, "bad tag in 'tags' - %s", err)
		}
	}
	return nil
}

// AttachTags implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) AttachTags(c context.Context, r *api.AttachTagsRequest) (resp *empty.Empty, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := validateMultiTagReq(r.Package, r.Instance, r.Tags); err != nil {
		return nil, err
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_WRITER); err != nil {
		return nil, err
	}

	// Actually attach the tags. This will also transactionally check the instance
	// exists and it has passed the processing successfully.
	inst := &model.Instance{
		InstanceID: model.ObjectRefToInstanceID(r.Instance),
		Package:    model.PackageKey(c, r.Package),
	}
	if err := model.AttachTags(c, inst, r.Tags, auth.CurrentIdentity(c)); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DetachTags implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) DetachTags(c context.Context, r *api.DetachTagsRequest) (resp *empty.Empty, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := validateMultiTagReq(r.Package, r.Instance, r.Tags); err != nil {
		return nil, err
	}

	// Check ACLs. Note this is scoped to OWNERS, see the proto doc.
	if _, err := impl.checkRole(c, r.Package, api.Role_OWNER); err != nil {
		return nil, err
	}

	// Verify the instance exists, per DetachTags contract.
	inst := &model.Instance{
		InstanceID: model.ObjectRefToInstanceID(r.Instance),
		Package:    model.PackageKey(c, r.Package),
	}
	if err := model.CheckInstanceExists(c, inst); err != nil {
		return nil, err
	}

	// Actually detach the tags.
	if err := model.DetachTags(c, inst, r.Tags); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

////////////////////////////////////////////////////////////////////////////////
// Version resolution and instance info fetching.

// ResolveVersion implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) ResolveVersion(c context.Context, r *api.ResolveVersionRequest) (resp *api.Instance, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if err := common.ValidateInstanceVersion(r.Version); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'version' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_READER); err != nil {
		return nil, err
	}

	// Actually resolve the version. This will return an appropriately grpc-tagged
	// error.
	inst, err := model.ResolveVersion(c, r.Package, r.Version)
	if err != nil {
		return nil, err
	}
	return inst.Proto(), nil
}

// GetInstanceURL implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) GetInstanceURL(c context.Context, r *api.GetInstanceURLRequest) (resp *api.ObjectURL, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	// Validate the request.
	if err := common.ValidatePackageName(r.Package); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'package' - %s", err)
	}
	if err := cas.ValidateObjectRef(r.Instance); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "bad 'instance' - %s", err)
	}

	// Check ACLs.
	if _, err := impl.checkRole(c, r.Package, api.Role_READER); err != nil {
		return nil, err
	}

	// Make sure this instance actually exists (without this check the caller
	// would be able to "probe" CAS namespace unrestricted).
	inst := (&model.Instance{}).FromProto(c, &api.Instance{
		Package:  r.Package,
		Instance: r.Instance,
	})
	if err := model.CheckInstanceExists(c, inst); err != nil {
		return nil, err
	}

	// Ask CAS generate an URL for us. Note that CAS does caching internally.
	return impl.cas.GetObjectURL(c, &api.GetObjectURLRequest{
		Object: r.Instance,
	})
}

// DescribeInstance implements the corresponding RPC method, see the proto doc.
func (impl *repoImpl) DescribeInstance(c context.Context, r *api.DescribeInstanceRequest) (resp *api.DescribeInstanceResponse, err error) {
	defer func() { err = grpcutil.GRPCifyAndLogErr(c, err) }()

	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
