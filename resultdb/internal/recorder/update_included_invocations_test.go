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

package recorder

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/resultdb/internal/span"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
	. "go.chromium.org/luci/resultdb/internal/testutil"
)

func TestValidateUpdateIncludedInvocationsRequest(t *testing.T) {
	t.Parallel()
	Convey(`TestValidateUpdateIncludedInvocationsRequest`, t, func() {
		Convey(`Valid`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldBeNil)
		})

		Convey(`Invalid including_invocation`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "x",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `including_invocation: does not match`)
		})
		Convey(`Invalid add_invocations`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"x"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `add_invocations: "x": does not match`)
		})
		Convey(`Invalid remove_invocations`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"x"},
			})
			So(err, ShouldErrLike, `remove_invocations: "x": does not match`)
		})
		Convey(`Include itself`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/a"},
				RemoveInvocations:   []string{"invocations/c"},
			})
			So(err, ShouldErrLike, `cannot include itself`)
		})
		Convey(`Add and remove same invocation`, func() {
			err := validateUpdateIncludedInvocationsRequest(&pb.UpdateIncludedInvocationsRequest{
				IncludingInvocation: "invocations/a",
				AddInvocations:      []string{"invocations/b"},
				RemoveInvocations:   []string{"invocations/b"},
			})
			So(err, ShouldErrLike, `cannot add and remove the same invocation(s) at the same time: ["invocations/b"]`)
		})
	})
}

func TestValidateIncludeRequest(t *testing.T) {
	Convey(`TestValidateIncludeRequest`, t, func() {
		Convey(`Valid`, func() {
			err := validateIncludeRequest(&pb.IncludeRequest{
				IncludingInvocation: "invocations/a",
				IncludedInvocation:  "invocations/b",
			})
			So(err, ShouldBeNil)
		})

		Convey(`Invalid including_invocation`, func() {
			err := validateIncludeRequest(&pb.IncludeRequest{
				IncludingInvocation: "x",
				IncludedInvocation:  "invocations/c",
			})
			So(err, ShouldErrLike, `including_invocation: does not match`)
		})

		Convey(`Invalid included_invocation`, func() {
			err := validateIncludeRequest(&pb.IncludeRequest{
				IncludingInvocation: "invocations/a",
				IncludedInvocation:  "x",
			})
			So(err, ShouldErrLike, `included_invocation: does not match`)
		})

		Convey(`Include itself`, func() {
			err := validateIncludeRequest(&pb.IncludeRequest{
				IncludingInvocation: "invocations/a",
				IncludedInvocation:  "invocations/a",
			})
			So(err, ShouldErrLike, `cannot include itself`)
		})
	})
}

func TestInclude(t *testing.T) {
	Convey(`TestInclude`, t, func() {
		ctx := SpannerTestContext(t)
		recorder := newTestRecorderServer()

		const token = "update token"
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(UpdateTokenMetadataKey, token))

		insInv := InsertInvocation

		assertIncluded := func(includedInvID span.InvocationID) {
			var throwAway span.InvocationID
			MustReadRow(ctx, "IncludedInvocations", span.InclusionKey("including", includedInvID), map[string]interface{}{
				"IncludedInvocationID": &throwAway,
			})
		}

		Convey(`Invalid request`, func() {
			_, err := recorder.Include(ctx, &pb.IncludeRequest{})
			So(err, ShouldHaveAppStatus, codes.InvalidArgument, `bad request: including_invocation: unspecified`)
		})

		req := &pb.IncludeRequest{
			IncludingInvocation: "invocations/including",
			IncludedInvocation:  "invocations/included",
		}

		Convey(`No including invocation`, func() {
			_, err := recorder.Include(ctx, req)
			So(err, ShouldHaveAppStatus, codes.NotFound, `invocations/including not found`)
		})

		Convey(`No included invocation`, func() {
			MustApply(ctx,
				insInv("including", pb.Invocation_ACTIVE, map[string]interface{}{"UpdateToken": token}),
			)
			_, err := recorder.Include(ctx, req)
			So(err, ShouldHaveAppStatus, codes.NotFound, `invocations/included not found`)
		})

		Convey(`Included invocation is active`, func() {
			MustApply(ctx,
				insInv("including", pb.Invocation_ACTIVE, map[string]interface{}{"UpdateToken": token}),
				insInv("included", pb.Invocation_ACTIVE, nil),
			)
			_, err := recorder.Include(ctx, req)
			So(err, ShouldBeNil)
		})

		Convey(`Idempotent`, func() {
			MustApply(ctx,
				insInv("including", pb.Invocation_ACTIVE, map[string]interface{}{"UpdateToken": token}),
				insInv("included", pb.Invocation_FINALIZED, nil),
			)

			_, err := recorder.Include(ctx, req)
			So(err, ShouldBeNil)

			_, err = recorder.Include(ctx, req)
			So(err, ShouldBeNil)
		})

		Convey(`Success`, func() {
			MustApply(ctx,
				insInv("including", pb.Invocation_ACTIVE, map[string]interface{}{"UpdateToken": token}),
				insInv("included", pb.Invocation_FINALIZED, nil),
			)

			_, err := recorder.Include(ctx, req)
			So(err, ShouldBeNil)
			assertIncluded("included")
		})
	})
}
