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
	"fmt"
	"math"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Min struct{}
type MinArgs struct {
	Input []float64 `pulumi:"input"`
}

type MinResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Min) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the smallest of the floats.")
}

func (*Min) Call(_ p.Context, args MinArgs) (MinResult, error) {
	if len(args.Input) == 0 {
		return MinResult{}, fmt.Errorf("input list must not be empty")
	}

	min := args.Input[0]
	for _, current := range args.Input {
		min = math.Min(min, current)
	}

	return MinResult{min}, nil
}
