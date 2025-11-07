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

type Contains struct{}
type ContainsArgs struct {
	Input   []interface{} `pulumi:"input"`
	Element interface{}   `pulumi:"element"`
}

type ContainsResult struct {
	Result bool `pulumi:"result"`
}

func (r *Contains) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns true if a list contains the given element and returns false otherwise.")
}

func (*Contains) Invoke(_ context.Context, input infer.FunctionRequest[ContainsArgs]) (infer.FunctionResponse[ContainsResult], error) {
	for _, element := range input.Input.Input {
		if jsonDeepEquals(element, input.Input.Element) {
			return infer.FunctionResponse[ContainsResult]{Output: ContainsResult{true}}, nil
		}
	}
	return infer.FunctionResponse[ContainsResult]{Output: ContainsResult{false}}, nil
}
