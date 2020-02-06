// Copyright 2020 The LUCI Authors.
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

package backend

import (
	"context"

	"cloud.google.com/go/spanner"

	"go.chromium.org/luci/common/logging"
	tsmoncommon "go.chromium.org/luci/common/tsmon"
	"go.chromium.org/luci/common/tsmon/metric"

	"go.chromium.org/luci/resultdb/internal/span"
)

var (
	expiredResultsDelayMetric = metric.NewInt(
		"resultdb/expired_results_delay",
		"How long overdue in seconds are the earliest results not yet purged",
		nil)
)

func recordExpiredResultsDelayMetric(ctx context.Context) {
	tsmoncommon.RegisterCallbackIn(ctx, func(ctx context.Context) {
		val, err := expiredResultsDelaySeconds(ctx)
		if err != nil {
			logging.Errorf(ctx, "Failed to get purge backlog delay: %s", err)
			return
		}
		expiredResultsDelayMetric.Set(ctx, val)
	})
}

// expiredResultsDelaySeconds computes the age of the oldest invocation
// pending to be purged in seconds.
func expiredResultsDelaySeconds(ctx context.Context) (int64, error) {
	st := spanner.NewStatement(`
		SELECT GREATEST(0, TIMESTAMP_DIFF(CURRENT_TIMESTAMP(), MIN(EarliestExpiration), SECOND)) AS BacklogSeconds
		FROM (
			SELECT (
				SELECT MIN(ExpectedTestResultsExpirationTime)
				FROM Invocations@{FORCE_INDEX=InvocationsByExpectedTestResultsExpiration}
				WHERE ShardId = TargetShard
				AND ExpectedTestResultsExpirationTime IS NOT NULL
			) AS EarliestExpiration
			FROM UNNEST(GENERATE_ARRAY(0, (
				SELECT MAX(ShardId)
				FROM Invocations@{FORCE_INDEX=InvocationsByExpectedTestResultsExpiration}
				WHERE ExpectedTestResultsExpirationTime IS NOT NULL
			))) AS TargetShard
		)
	`)
	var ret int64
	if err := span.QueryFirstRow(ctx, span.Client(ctx).Single(), st, &ret); err != nil {
		return 0, err
	}
	return ret, nil
}
