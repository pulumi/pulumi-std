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

type Tobool struct{}
type ToboolArgs struct {
	Input interface{} `pulumi:"input"`
}

type ToboolResult struct {
	Result *bool `pulumi:"result,optional"`
}

func (r *Tobool) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a boolean value. Only boolean values, null, and the exact strings
	"true" and "false" can be converted to boolean. All other values will result in an error.`)
}

func (*Tobool) Invoke(
	_ context.Context,
	input infer.FunctionRequest[ToboolArgs],
) (infer.FunctionResponse[ToboolResult], error) {
	if input.Input.Input == nil {
		return infer.FunctionResponse[ToboolResult]{Output: ToboolResult{nil}}, nil
	}
	if v, ok := input.Input.Input.(bool); ok {
		return infer.FunctionResponse[ToboolResult]{Output: ToboolResult{toBoolPtr(v)}}, nil
	} else if str, ok := input.Input.Input.(string); ok {
		v, err := strconv.ParseBool(str)
		if err != nil {
			return infer.FunctionResponse[ToboolResult]{Output: ToboolResult{nil}}, fmt.Errorf("%v is not a boolean value", input.Input.Input)
		}
		return infer.FunctionResponse[ToboolResult]{Output: ToboolResult{toBoolPtr(v)}}, nil
	}

	return infer.FunctionResponse[ToboolResult]{
			Output: ToboolResult{nil},
		}, fmt.Errorf(
			"%v is not a boolean value",
			input.Input.Input,
		)
}

func toBoolPtr(b bool) *bool {
	return &b
}
