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

type Max struct{}
type MaxArgs struct {
	Input []float64 `pulumi:"input"`
}

type MaxResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Max) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the largest of the floats.")
}

func (*Max) Call(_ p.Context, args MaxArgs) (MaxResult, error) {
	if len(args.Input) == 0 {
		return MaxResult{}, fmt.Errorf("input list must not be empty")
	}

	max := args.Input[0]
	for _, current := range args.Input {
		max = math.Max(max, current)
	}

	return MaxResult{max}, nil
}
