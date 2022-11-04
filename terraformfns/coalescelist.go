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

package terraformfns

import (
	"errors"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Coalescelist struct{}
type CoalescelistArgs struct {
	Input [][]interface{} `pulumi:"input"`
}

type CoalescelistResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Coalescelist) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the first non-empty list from the given list of lists.")
}

func (*Coalescelist) Call(ctx p.Context, args CoalescelistArgs) (CoalescelistResult, error) {
	for _, list := range args.Input {
		if len(list) > 0 {
			return CoalescelistResult{list}, nil
		}
	}

	return CoalescelistResult{}, errors.New("no non-empty list found")
}
