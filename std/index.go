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
	"errors"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Index struct{}
type IndexArgs struct {
	Input   []interface{} `pulumi:"input"`
	Element interface{}   `pulumi:"element"`
}

type IndexResult struct {
	Result int `pulumi:"result"`
}

func (r *Index) Annotate(a infer.Annotator) {
	a.Describe(r, "Finds the index of a given element in a list.")
}

func (*Index) Call(_ p.Context, args IndexArgs) (IndexResult, error) {
	if len(args.Input) == 0 {
		return IndexResult{}, errors.New("input list must not be empty when using the index function")
	}

	for index, value := range args.Input {
		if jsonDeepEquals(value, args.Element) {
			return IndexResult{index}, nil
		}
	}

	return IndexResult{}, errors.New("item not found")
}
