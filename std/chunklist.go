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
	"errors"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Chunklist struct{}
type ChunklistArgs struct {
	Input []interface{} `pulumi:"input"`
	Size  int           `pulumi:"size"`
}

type ChunklistResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Chunklist) Annotate(a infer.Annotator) {
	a.Describe(r, "Splits a single list into multiple lists where each has at most the given number of elements.")
}

func (*Chunklist) Invoke(_ context.Context, input infer.FunctionRequest[ChunklistArgs]) (infer.FunctionResponse[ChunklistResult], error) {
	if input.Input.Size < 0 {
		return infer.FunctionResponse[ChunklistResult]{Output: ChunklistResult{}}, errors.New("size must be greater than or equal to 0")
	}

	output := make([]interface{}, 0)
	for i := 0; i < len(input.Input.Input); i += input.Input.Size {
		end := i + input.Input.Size
		if end > len(input.Input.Input) {
			end = len(input.Input.Input)
		}
		output = append(output, input.Input.Input[i:end])
	}

	return infer.FunctionResponse[ChunklistResult]{Output: ChunklistResult{output}}, nil
}
