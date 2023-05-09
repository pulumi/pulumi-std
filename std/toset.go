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
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Toset struct{}
type TosetArgs struct {
	Input []interface{} `pulumi:"input"`
}

type TosetResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Toset) Annotate(a infer.Annotator) {
	a.Describe(r, "Converts its argument to a set value.")
}

func (*Toset) Call(_ p.Context, args TosetArgs) (TosetResult, error) {
	if len(args.Input) == 0 {
		return TosetResult{make([]interface{}, 0)}, nil
	}

	return TosetResult{convertListSameType(removeDuplicateElements(args.Input))}, nil

}

func removeDuplicateElements(input []interface{}) []interface{} {
	keys := make(map[interface{}]bool)
	res := make([]interface{}, 0)

	for _, e := range input {
		if _, value := keys[e]; !value {
			keys[e] = true
			res = append(res, e)
		}
	}
	return res
}
