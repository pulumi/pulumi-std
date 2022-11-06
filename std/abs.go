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
	"math"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Abs struct{}
type AbsArgs struct {
	Input float64 `pulumi:"input"`
}

type AbsResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Abs) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns the absolute value of a given float. 
Example: abs(1) returns 1, and abs(-1) would also return 1, whereas abs(-3.14) would return 3.14.`)
}

func (*Abs) Call(ctx p.Context, input AbsArgs) (AbsResult, error) {
	return AbsResult{math.Abs(input.Input)}, nil
}
