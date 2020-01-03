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

package span

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/golang/protobuf/ptypes"
	durpb "github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/grpcutil"

	"go.chromium.org/luci/resultdb/internal/pagination"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

// MustParseTestResultName retrieves the invocation ID, unescaped test id, and
// result ID.
// Panics if the name is invalid. Useful for situations when name was already
// validated.
func MustParseTestResultName(name string) (invID InvocationID, testID, resultID string) {
	invIDStr, testID, resultID, err := pbutil.ParseTestResultName(name)
	if err != nil {
		panic(err)
	}
	invID = InvocationID(invIDStr)
	return
}

// ReadTestResult reads specified TestResult within the transaction.
// If the TestResult does not exist, the returned error is annotated with
// NotFound GRPC code.
func ReadTestResult(ctx context.Context, txn Txn, name string) (*pb.TestResult, error) {
	invID, testID, resultID := MustParseTestResultName(name)
	tr := &pb.TestResult{
		Name:     name,
		TestId:   testID,
		ResultId: resultID,
		Expected: true,
	}

	var maybeUnexpected spanner.NullBool
	var micros int64
	var summaryHTML Compressed
	err := ReadRow(ctx, txn, "TestResults", invID.Key(testID, resultID), map[string]interface{}{
		"Variant":         &tr.Variant,
		"IsUnexpected":    &maybeUnexpected,
		"Status":          &tr.Status,
		"SummaryHTML":     &summaryHTML,
		"StartTime":       &tr.StartTime,
		"RunDurationUsec": &micros,
		"Tags":            &tr.Tags,
		"InputArtifacts":  &tr.InputArtifacts,
		"OutputArtifacts": &tr.OutputArtifacts,
	})
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		return nil, errors.Reason("%q not found", name).
			InternalReason("%s", err).
			Tag(grpcutil.NotFoundTag).
			Err()
	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", name).Err()
	}

	tr.SummaryHtml = string(summaryHTML)
	populateExpectedField(tr, maybeUnexpected)
	populateDurationField(tr, micros)
	return tr, nil
}

// TestResultQuery specifies test results to fetch.
type TestResultQuery struct {
	InvocationIDs InvocationIDSet
	Predicate     *pb.TestResultPredicate // Predicate.Invocation must be nil.
	PageSize      int                     // must be positive
	PageToken     string
}

func queryTestResults(ctx context.Context, txn *spanner.ReadOnlyTransaction, q TestResultQuery, f func(tr *pb.TestResult) error) (err error) {
	switch {
	case q.PageSize < 0:
		panic("PageSize < 0")
	}

	from := "TestResults tr"
	if q.Predicate.GetExpectancy() == pb.TestResultPredicate_VARIANTS_WITH_UNEXPECTED_RESULTS {
		// We must return only test results of test variants that have unexpected results.
		//
		// The following query ensures that we first select test variants with
		// unexpected results, and then for each variant do a lookup in TestResults
		// table.
		from = `
			VariantsWithUnexpectedResults vur
			JOIN@{FORCE_JOIN_ORDER=TRUE} TestResults tr
				ON vur.TestId = tr.TestId AND vur.VariantHash = tr.VariantHash
		`
	}

	limit := ""
	if q.PageSize > 0 {
		limit = `LIMIT @limit`
	}

	st := spanner.NewStatement(fmt.Sprintf(`
		WITH VariantsWithUnexpectedResults AS (
			# Note: this query is not executed if it ends up not used in the top-level
			# query.
			SELECT DISTINCT TestId, VariantHash
			FROM TestResults@{FORCE_INDEX=UnexpectedTestResults}
			WHERE IsUnexpected AND InvocationId IN UNNEST(@invIDs)
		)
		SELECT
			tr.InvocationId,
			tr.TestId,
			tr.ResultId,
			tr.Variant,
			tr.IsUnexpected,
			tr.Status,
			tr.SummaryHtml,
			tr.StartTime,
			tr.RunDurationUsec,
			tr.Tags,
			tr.InputArtifacts,
			tr.OutputArtifacts
		FROM %s
		WHERE InvocationId IN UNNEST(@invIDs)
			# Skip test results after the one specified in the page token.
			AND (
				(tr.InvocationId > @afterInvocationId) OR
				(tr.InvocationId = @afterInvocationId AND tr.TestId > @afterTestId) OR
				(tr.InvocationId = @afterInvocationId AND tr.TestId = @afterTestId AND tr.ResultId > @afterResultId)
			)
			AND REGEXP_CONTAINS(tr.TestId, @TestIdRegexp)
		ORDER BY tr.InvocationId, tr.TestId, tr.ResultId
		%s
	`, from, limit))
	st.Params["invIDs"] = q.InvocationIDs
	st.Params["limit"] = q.PageSize

	testIDRegexp := q.Predicate.GetTestIdRegexp()
	if testIDRegexp == "" {
		testIDRegexp = ".*"
	}
	st.Params["TestIdRegexp"] = fmt.Sprintf("^%s$", testIDRegexp)

	st.Params["afterInvocationId"],
		st.Params["afterTestId"],
		st.Params["afterResultId"],
		err = parseTestObjectPageToken(q.PageToken)
	if err != nil {
		return
	}

	if q.Predicate.GetVariant() != nil {
		// TODO(nodir): add support for q.Predicate.Variant.
		return grpcutil.Unimplemented
	}

	var summaryHTML Compressed
	var b Buffer
	return query(ctx, txn, st, func(row *spanner.Row) error {
		var invID InvocationID
		var maybeUnexpected spanner.NullBool
		var micros int64
		tr := &pb.TestResult{}
		err = b.FromSpanner(row,
			&invID,
			&tr.TestId,
			&tr.ResultId,
			&tr.Variant,
			&maybeUnexpected,
			&tr.Status,
			&summaryHTML,
			&tr.StartTime,
			&micros,
			&tr.Tags,
			&tr.InputArtifacts,
			&tr.OutputArtifacts,
		)
		if err != nil {
			return err
		}

		tr.Name = pbutil.TestResultName(string(invID), tr.TestId, tr.ResultId)
		tr.SummaryHtml = string(summaryHTML)
		populateExpectedField(tr, maybeUnexpected)
		populateDurationField(tr, micros)

		return f(tr)
	})
}

// QueryTestResults reads test results matching the predicate.
// Returned test results from the same invocation are contiguous.
func QueryTestResults(ctx context.Context, txn *spanner.ReadOnlyTransaction, q TestResultQuery) (trs []*pb.TestResult, nextPageToken string, err error) {
	switch {
	case q.PageSize <= 0:
		panic("PageSize <= 0")
	}

	trs = make([]*pb.TestResult, 0, q.PageSize)
	err = queryTestResults(ctx, txn, q, func(tr *pb.TestResult) error {
		trs = append(trs, tr)
		return nil
	})
	if err != nil {
		trs = nil
		return
	}

	// If we got pageSize results, then we haven't exhausted the collection and
	// need to return the next page token.
	if len(trs) == q.PageSize {
		last := trs[q.PageSize-1]
		invID, testID, resultID := MustParseTestResultName(last.Name)
		nextPageToken = pagination.Token(string(invID), testID, resultID)
	}
	return
}

func populateDurationField(tr *pb.TestResult, micros int64) {
	tr.Duration = FromMicros(micros)
}

func populateExpectedField(tr *pb.TestResult, maybeUnexpected spanner.NullBool) {
	tr.Expected = !maybeUnexpected.Valid || !maybeUnexpected.Bool
}

// ToMicros converts a duration.Duration proto to microseconds.
func ToMicros(d *durpb.Duration) int64 {
	if d == nil {
		return 0
	}
	return 1e6*d.Seconds + int64(1e-3*float64(d.Nanos))
}

// FromMicros converts microseconds to a duration.Duration proto.
func FromMicros(micros int64) *durpb.Duration {
	return ptypes.DurationProto(time.Duration(1e3 * micros))
}

// parseTestObjectPageToken parses the page token into invocation ID, test id
// and a test object id.
func parseTestObjectPageToken(pageToken string) (inv InvocationID, testID, objID string, err error) {
	switch pos, tokErr := pagination.ParseToken(pageToken); {
	case tokErr != nil:
		err = encapsulatePageTokenError(tokErr)

	case pos == nil:

	case len(pos) != 3:
		err = encapsulatePageTokenError(errors.Reason("expected 3 position strings, got %q", pos).Err())

	default:
		inv = InvocationID(pos[0])
		testID = pos[1]
		objID = pos[2]
	}

	return
}

// encapsulatePageTokenError returns a generic error message that a page token
// is invalid and records err as an internal error.
// The returned error is anontated with INVALID_ARUGMENT code.
func encapsulatePageTokenError(err error) error {
	return errors.Reason("invalid page_token").InternalReason("%s", err).Tag(grpcutil.InvalidArgumentTag).Err()
}
