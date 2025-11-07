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

func (*Contains) Call(_ context.Context, args ContainsArgs) (ContainsResult, error) {
	for _, element := range args.Input {
		if jsonDeepEquals(element, args.Element) {
			return ContainsResult{true}, nil
		}
	}
	return ContainsResult{false}, nil
}
