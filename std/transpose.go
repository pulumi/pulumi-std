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
	"sort"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Transpose struct{}
type TransposeArgs struct {
	Input map[string][]string `pulumi:"input"`
}

type TransposeResult struct {
	Result map[string][]string `pulumi:"result"`
}

func (r *Transpose) Annotate(a infer.Annotator) {
	a.Describe(r, `Takes a map of lists of strings and swaps the keys and values to return a new map of lists of strings.`)
}

func (*Transpose) Call(_ p.Context, args TransposeArgs) (TransposeResult, error) {
	res := make(map[string][]string)
	for k, l := range args.Input {
		for _, v := range l {
			if _, ok := res[v]; ok {
				res[v] = append(res[v], k)
			} else {
				res[v] = []string{k}
			}
		}
	}
	// sort result to produce an expected ordering
	for _, v := range res {
		sort.Strings(v)
	}
	return TransposeResult{res}, nil
}
