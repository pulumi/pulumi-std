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

type Values struct{}
type ValuesArgs struct {
	Input map[string]interface{} `pulumi:"input"`
}

type ValuesResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Values) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a list of the values of the map.")
}

func (*Values) Call(_ p.Context, args ValuesArgs) (ValuesResult, error) {
	keys := make([]string, 0, len(args.Input))
	for key := range args.Input {
		keys = append(keys, key)
	}
	// Sort the keys so that the output is deterministic.
	sort.Strings(keys)

	values := make([]interface{}, 0, len(args.Input))
	for _, key := range keys {
		values = append(values, args.Input[key])
	}
	return ValuesResult{values}, nil
}
