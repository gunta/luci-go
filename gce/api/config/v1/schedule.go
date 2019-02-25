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

package config

import (
	"go.chromium.org/luci/config/validation"
)

// Validate validates this schedule.
func (s *Schedule) Validate(c *validation.Context) {
	if s.GetAmount() < 0 {
		c.Errorf("amount must be non-negative")
	}
	c.Enter("length")
	s.GetLength().Validate(c)
	switch n, err := s.Length.ToSeconds(); {
	case err != nil:
		c.Errorf("%s", err)
	case n == 0:
		c.Errorf("duration or seconds is required")
	}
	c.Exit()
	c.Enter("start")
	s.GetStart().Validate(c)
	c.Exit()
}