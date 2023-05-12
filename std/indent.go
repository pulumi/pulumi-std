// Copyright 2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package std

import (
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Indent struct{}
type IndentArgs struct {
	Input  string `pulumi:"input"`
	Spaces int    `pulumi:"spaces"`
}

type IndentResult struct {
	Result string `pulumi:"result"`
}

func (r *Indent) Annotate(a infer.Annotator) {
	a.Describe(r, "Adds a given number of spaces after each newline character in the given string.")
}

func (*Indent) Call(_ p.Context, args IndentArgs) (IndentResult, error) {
	pad := strings.Repeat(" ", args.Spaces)
	return IndentResult{strings.Replace(args.Input, "\n", "\n"+pad, -1)}, nil
}
