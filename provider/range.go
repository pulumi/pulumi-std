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
	"context"
	"math"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Range struct{}
type RangeArgs struct {
	Limit float64  `pulumi:"limit"`
	Start *float64 `pulumi:"start,optional"`
	Step  *float64 `pulumi:"step,optional"`
}

type RangeResults struct {
	Result []float64 `pulumi:"result"`
}

func (r *Range) Annotate(a infer.Annotator) {
	a.Describe(r, `Generates a list of numbers using a start value, a limit value, and a step value.
Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
depending on whether limit is greater than or less than start.`)
}

func genRange(start, limit, step float64) []float64 {
	if start == limit {
		return []float64{}
	}
	if step == 0 {
		if start < limit {
			step = 1
		} else {
			step = -1
		}
	} else if (start < limit && step < 0) || (start > limit && step > 0) {
		return []float64{}
	}
	length := int(math.Ceil((limit - start) / step))
	res := make([]float64, length)
	for i := 0; i < length; i++ {
		res[i] = start + (float64(i) * step)
	}
	return res
}

func (*Range) Invoke(
	_ context.Context,
	input infer.FunctionRequest[RangeArgs],
) (infer.FunctionResponse[RangeResults], error) {
	if input.Input.Start == nil {
		start := 0.0
		input.Input.Start = &start
	}

	if input.Input.Step == nil {
		step := 1.0
		if *input.Input.Start > input.Input.Limit {
			step = -1.0
		}

		input.Input.Step = &step
	}

	return infer.FunctionResponse[RangeResults]{
		Output: RangeResults{genRange(*input.Input.Start, input.Input.Limit, *input.Input.Step)},
	}, nil
}
