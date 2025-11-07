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
	"context"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Trimprefix struct{}
type TrimprefixArgs struct {
	Input  string `pulumi:"input"`
	Prefix string `pulumi:"prefix"`
}

type TrimprefixResult struct {
	Result string `pulumi:"result"`
}

func (r *Trimprefix) Annotate(a infer.Annotator) {
	a.Describe(r, `Removes the specified prefix from the start of the given string, if present.`)
}

func (*Trimprefix) Invoke(_ context.Context, input infer.FunctionRequest[TrimprefixArgs]) (infer.FunctionResponse[TrimprefixResult], error) {
	return infer.FunctionResponse[TrimprefixResult]{Output: TrimprefixResult{strings.TrimPrefix(input.Input.Input, input.Input.Prefix)}}, nil
}
