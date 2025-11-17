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

type Matchkeys struct{}
type MatchkeysArgs struct {
	Values     []string `pulumi:"values"`
	keys       []string `pulumi:"keys"`
	SearchList []string `pulumi:"searchList"`
}

type MatchkeysResult struct {
	Result []string `pulumi:"result"`
}

func (r *Matchkeys) Annotate(a infer.Annotator) {
	a.Describe(r, `For two lists values and keys of equal length,
returns all elements from values where the corresponding element from keys exists in the searchset list.`)
}

func (*Matchkeys) Invoke(
	_ context.Context,
	input infer.FunctionRequest[MatchkeysArgs],
) (infer.FunctionResponse[MatchkeysResult], error) {
	if len(input.Input.Values) != len(input.Input.keys) {
		return infer.FunctionResponse[MatchkeysResult]{
				Output: MatchkeysResult{},
			}, errors.New(
				"values and keys must be of equal length",
			)
	}

	var output []string
	for i, key := range input.Input.keys {
		for _, search := range input.Input.SearchList {
			if key == search {
				output = append(output, input.Input.Values[i])
			}
		}
	}

	return infer.FunctionResponse[MatchkeysResult]{Output: MatchkeysResult{output}}, nil
}
