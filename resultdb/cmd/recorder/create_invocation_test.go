// Copyright 2019 The LUCI Authors.
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

package main

import (
	"testing"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/testing/prpctest"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/testutil"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateInvocationDeadline(t *testing.T) {
	Convey(`ValidateInvocationDeadline`, t, func() {
		now := testclock.TestRecentTimeUTC

		Convey(`deadline in the past`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(-time.Hour))
			err := validateInvocationDeadline(deadline, now)
			So(err, ShouldErrLike, `must be at least 10 seconds in the future`)
		})

		Convey(`deadline 5s in the future`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(5 * time.Second))
			err := validateInvocationDeadline(deadline, now)
			So(err, ShouldErrLike, `must be at least 10 seconds in the future`)
		})

		Convey(`deadline in the future`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(1e3 * time.Hour))
			err := validateInvocationDeadline(deadline, now)
			So(err, ShouldErrLike, `must be before 48h in the future`)
		})
	})
}

func TestValidateCreateInvocationRequest(t *testing.T) {
	t.Parallel()
	now := testclock.TestRecentTimeUTC
	Convey(`TestValidateCreateInvocationRequest`, t, func() {
		Convey(`empty`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{}, now)
			So(err, ShouldErrLike, `invocation_id: unspecified`)
		})

		Convey(`invalid id`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "1",
			}, now)
			So(err, ShouldErrLike, `invocation_id: does not match`)
		})

		Convey(`reserved prefix`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "build-1",
			}, now)
			So(err, ShouldErrLike, `must have id starting with "u:"`)
		})

		Convey(`invalid request id`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "u:a",
				RequestId:    "😃",
			}, now)
			So(err, ShouldErrLike, "request_id: does not match")
		})

		Convey(`invalid tags`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "u:abc",
				Invocation: &pb.Invocation{
					Tags: pbutil.StringPairs("1", "a"),
				},
			}, now)
			So(err, ShouldErrLike, `invocation.tags: "1":"a": key: does not match`)
		})

		Convey(`invalid deadline`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(-time.Hour))
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "u:abc",
				Invocation: &pb.Invocation{
					Deadline: deadline,
				},
			}, now)
			So(err, ShouldErrLike, `invocation: deadline: must be at least 10 seconds in the future`)
		})

		Convey(`invalid variant def`, func() {
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "u:abc",
				Invocation: &pb.Invocation{
					BaseTestVariantDef: &pb.VariantDef{
						Def: map[string]string{"1": "a"},
					},
				},
			}, now)
			So(err, ShouldErrLike, `invocation.base_test_variant_def: "1":"a": key: does not match`)
		})

		Convey(`valid`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(time.Hour))
			err := validateCreateInvocationRequest(&pb.CreateInvocationRequest{
				InvocationId: "u:abc",
				Invocation: &pb.Invocation{
					Deadline: deadline,
					Tags:     pbutil.StringPairs("a", "b", "a", "c", "d", "e"),
					BaseTestVariantDef: &pb.VariantDef{
						Def: map[string]string{
							"a": "b",
							"c": "d",
						},
					},
				},
			}, now)
			So(err, ShouldBeNil)
		})
	})
}

func TestCreateInvocation(t *testing.T) {
	Convey(`TestCreateInvocation`, t, func() {
		ctx := testutil.SpannerTestContext(t)

		// Init auth state.
		ctx = authtest.MockAuthConfig(ctx)
		authState := &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
		}
		ctx = auth.WithState(ctx, authState)

		// Mock time.
		now := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
		ctx, _ = testclock.UseTime(ctx, now)

		// Setup a full HTTP server in order to retrieve response headers.
		server := &prpctest.Server{}
		pb.RegisterRecorderServer(server, NewRecorderServer())
		server.Start(ctx)
		defer server.Close()
		client, err := server.NewClient()
		So(err, ShouldBeNil)
		recorder := pb.NewRecorderPRPCClient(client)

		Convey(`permission denied`, func() {
			_, err := recorder.CreateInvocation(ctx, &pb.CreateInvocationRequest{InvocationId: "x"})
			So(err, ShouldErrLike, `anonymous:anonymous is not allowed to create invocations`)
			So(grpcutil.Code(err), ShouldEqual, codes.PermissionDenied)
		})

		authState.IdentityGroups = []string{createInvocationGroup}

		Convey(`empty request`, func() {
			_, err := recorder.CreateInvocation(ctx, &pb.CreateInvocationRequest{})
			So(err, ShouldErrLike, `bad request: invocation_id: unspecified`)
			So(grpcutil.Code(err), ShouldEqual, codes.InvalidArgument)
		})

		req := &pb.CreateInvocationRequest{
			InvocationId: "u:inv",
			Invocation:   &pb.Invocation{},
		}

		Convey(`already exists`, func() {
			_, err := span.Client(ctx).Apply(ctx, []*spanner.Mutation{
				testutil.InsertInvocation(req.InvocationId, 1, "", testclock.TestRecentTimeUTC),
			})
			So(err, ShouldBeNil)

			_, err = recorder.CreateInvocation(ctx, req)
			So(err, ShouldErrLike, `already exists`)
			So(grpcutil.Code(err), ShouldEqual, codes.AlreadyExists)
		})

		Convey(`unsorted tags`, func() {
			req.Invocation.Tags = pbutil.StringPairs("b", "2", "a", "1")
			inv, err := recorder.CreateInvocation(ctx, req)
			So(err, ShouldBeNil)
			So(inv.Tags, ShouldResemble, pbutil.StringPairs("a", "1", "b", "2"))

		})

		Convey(`no invocation in request`, func() {
			inv, err := recorder.CreateInvocation(ctx, &pb.CreateInvocationRequest{InvocationId: "u:inv"})
			So(err, ShouldBeNil)
			So(inv.Name, ShouldEqual, "invocations/u:inv")
		})

		Convey(`idempotent`, func() {
			req := &pb.CreateInvocationRequest{
				InvocationId: "u:inv",
				Invocation:   &pb.Invocation{},
				RequestId:    "request id",
			}
			res, err := recorder.CreateInvocation(ctx, req)
			So(err, ShouldBeNil)

			res2, err := recorder.CreateInvocation(ctx, req)
			So(err, ShouldBeNil)
			So(res2, ShouldResembleProto, res)
		})

		Convey(`end to end`, func() {
			deadline := pbutil.MustTimestampProto(now.Add(time.Hour))
			headers := &metadata.MD{}
			req := &pb.CreateInvocationRequest{
				InvocationId: "u:inv",
				Invocation: &pb.Invocation{
					Deadline: deadline,
					Tags:     pbutil.StringPairs("a", "1", "b", "2"),
					BaseTestVariantDef: &pb.VariantDef{
						Def: map[string]string{
							"bucket":  "ci",
							"builder": "linux-rel",
						},
					},
				},
			}
			inv, err := recorder.CreateInvocation(ctx, req, prpc.Header(headers))
			So(err, ShouldBeNil)

			expected := proto.Clone(req.Invocation).(*pb.Invocation)
			proto.Merge(expected, &pb.Invocation{
				Name:       "invocations/u:inv",
				State:      pb.Invocation_ACTIVE,
				CreateTime: pbutil.MustTimestampProto(now),
			})
			So(inv, ShouldResembleProto, expected)

			So(headers.Get(updateTokenMetadataKey), ShouldHaveLength, 1)

			txn, err := span.Client(ctx).BatchReadOnlyTransaction(ctx, spanner.StrongRead())
			So(err, ShouldBeNil)
			defer txn.Close()

			inv, err = span.ReadInvocationFull(ctx, txn, "u:inv")
			So(err, ShouldBeNil)
			So(inv, ShouldResembleProto, expected)
		})
	})
}