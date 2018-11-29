// Copyright 2018 The LUCI Authors.
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

// Command lucicfg is CLI for LUCI config generator.
package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/lucicfg"
	"go.chromium.org/luci/starlark/interpreter"
)

// It's just a tiny example for now.

const input = `
load("@proto//luci/config/project_config.proto", config_pb="config")

def gen(ctx):
	ctx.config_set['project.cfg'] = config_pb.ProjectCfg(
			name = 'test',
			access = ['group:all'],
	)
core.generator(impl = gen)
`

func main() {
	ctx := context.Background()

	state, err := lucicfg.Generate(ctx, lucicfg.Inputs{
		Code: interpreter.MemoryLoader(map[string]string{
			"main.star": input,
		}),
		Entry: "main.star",
	})

	// Report all encountered errors, with stack traces.
	if err != nil {
		errors.WalkLeaves(err, func(err error) bool {
			if bt, ok := err.(lucicfg.BacktracableError); ok {
				fmt.Fprintf(os.Stderr, "%s\n\n", bt.Backtrace())
			} else {
				fmt.Fprintf(os.Stderr, "%s\n\n", err)
			}
			return true
		})
		os.Exit(1)
	}

	files := make([]string, 0, len(state.Configs))
	for f := range state.Configs {
		files = append(files, f)
	}
	sort.Strings(files)

	for _, f := range files {
		fmt.Println("--------------------------------------------------")
		fmt.Println(f)
		fmt.Println("--------------------------------------------------")
		fmt.Print(state.Configs[f])
		fmt.Println("--------------------------------------------------")
	}
}