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

type Startswith struct{}
type StartswithArgs struct {
	Input  string `pulumi:"input"`
	Prefix string `pulumi:"prefix"`
}

type StartswithResult struct {
	Result bool `pulumi:"result"`
}

func (r *Startswith) Annotate(a infer.Annotator) {
	a.Describe(r, "Determines if the input string starts with the suffix.")
}

func (*Startswith) Call(_ p.Context, args StartswithArgs) (StartswithResult, error) {
	return StartswithResult{strings.HasPrefix(args.Input, args.Prefix)}, nil
}
