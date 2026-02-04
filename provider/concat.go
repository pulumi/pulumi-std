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

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Concat struct{}
type ConcatArgs struct {
	Input [][]interface{} `pulumi:"input"`
}

type ConcatResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Concat) Annotate(a infer.Annotator) {
	a.Describe(r, "Combines two or more lists into a single list.")
}

func (*Concat) Invoke(
	_ context.Context,
	input infer.FunctionRequest[ConcatArgs],
) (infer.FunctionResponse[ConcatResult], error) {
	output := make([]interface{}, 0)
	for _, list := range input.Input.Input {
		output = append(output, list...)
	}

	return infer.FunctionResponse[ConcatResult]{Output: ConcatResult{output}}, nil
}
