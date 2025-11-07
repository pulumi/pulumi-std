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
	"fmt"
	"strconv"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Tonumber struct{}
type TonumberArgs struct {
	Input interface{} `pulumi:"input"`
}

type TonumberResult struct {
	Result *float64 `pulumi:"result,optional"`
}

func (r *Tonumber) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a number value. Only number values, null, and strings
	containing decimal representations of numbers can be converted to number. All other values will result in an error`)
}

func (*Tonumber) Invoke(_ context.Context, input infer.FunctionRequest[TonumberArgs]) (infer.FunctionResponse[TonumberResult], error) {
	if input.Input.Input == nil {
		return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{nil}}, nil
	}
	if v, ok := input.Input.Input.(int); ok {
		return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{toFloat64Ptr(float64(v))}}, nil
	} else if v, ok := input.Input.Input.(float64); ok {
		return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{toFloat64Ptr(v)}}, nil
	} else if str, ok := input.Input.Input.(string); ok {
		vInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			vFloat, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{nil}}, fmt.Errorf("%v is not a number value", input.Input.Input)
			}
			return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{toFloat64Ptr(vFloat)}}, nil
		}
		return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{toFloat64Ptr(float64(vInt))}}, nil
	}
	return infer.FunctionResponse[TonumberResult]{Output: TonumberResult{nil}}, fmt.Errorf("%v is not a number value", input.Input.Input)
}

func toFloat64Ptr(f float64) *float64 {
	return &f
}
