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

type Flatten struct{}
type FlattenArgs struct {
	Input []interface{} `pulumi:"input"`
}

type FlattenResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Flatten) Annotate(a infer.Annotator) {
	a.Describe(r, `Flattens lists of lists down to a flat list of primitive values,
eliminating any nested lists recursively.`)
}

func traverse(input []interface{}, elem func(interface{})) {
	for _, value := range input {
		switch value := value.(type) {
		case []interface{}:
			traverse(value, elem)
		default:
			elem(value)
		}
	}
}

func flatten(input []interface{}) []interface{} {
	output := make([]interface{}, 0)
	traverse(input, func(value interface{}) {
		output = append(output, value)
	})

	return output
}

func (*Flatten) Invoke(_ context.Context, input infer.FunctionRequest[FlattenArgs]) (infer.FunctionResponse[FlattenResult], error) {
	return infer.FunctionResponse[FlattenResult]{Output: FlattenResult{flatten(input.Input.Input)}}, nil
}
