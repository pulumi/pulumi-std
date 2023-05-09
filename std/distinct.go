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

type Distinct struct{}
type DistinctArgs struct {
	Input []interface{} `pulumi:"input"`
}

type DistinctResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Distinct) Annotate(a infer.Annotator) {
	a.Describe(r, "Removes duplicate items from a list.")
}

func (*Distinct) Call(_ p.Context, args DistinctArgs) (DistinctResult, error) {
	seen := make(map[interface{}]bool)
	output := make([]interface{}, 0)
	for _, value := range args.Input {
		if !seen[value] {
			seen[value] = true
			output = append(output, value)
		}
	}

	return DistinctResult{output}, nil
}
