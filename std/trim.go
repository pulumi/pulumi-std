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

type Trim struct{}
type TrimArgs struct {
	Input  string `pulumi:"input"`
	CutSet string `pulumi:"cutset"`
}

type TrimResult struct {
	Result string `pulumi:"result"`
}

func (r *Trim) Annotate(a infer.Annotator) {
	a.Describe(r, `Removes the specified set of characters from the start and end of the given string.`)
}

func (*Trim) Invoke(_ context.Context, input infer.FunctionRequest[TrimArgs]) (infer.FunctionResponse[TrimResult], error) {
	return infer.FunctionResponse[TrimResult]{Output: TrimResult{strings.Trim(input.Input.Input, input.Input.CutSet)}}, nil
}
