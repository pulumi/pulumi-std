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

func (*Matchkeys) Call(_ p.Context, args MatchkeysArgs) (MatchkeysResult, error) {
	if len(args.Values) != len(args.keys) {
		return MatchkeysResult{}, errors.New("values and keys must be of equal length")
	}

	var output []string
	for i, key := range args.keys {
		for _, search := range args.SearchList {
			if key == search {
				output = append(output, args.Values[i])
			}
		}
	}

	return MatchkeysResult{output}, nil
}
