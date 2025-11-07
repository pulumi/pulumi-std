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

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Tostring struct{}
type TostringArgs struct {
	Input interface{} `pulumi:"input"`
}

type TostringResult struct {
	Result *string `pulumi:"result,optional"`
}

func (r *Tostring) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a string value. Only primitive types (string, number, bool)
	and null can be converted to string. All other values will result in an error.`)
}

func (*Tostring) Invoke(_ context.Context, input infer.FunctionRequest[TostringArgs]) (infer.FunctionResponse[TostringResult], error) {
	if input.Input.Input == nil {
		return infer.FunctionResponse[TostringResult]{Output: TostringResult{nil}}, nil
	}
	if v, ok := input.Input.Input.(string); ok {
		return infer.FunctionResponse[TostringResult]{Output: TostringResult{toStrPtr(v)}}, nil
	} else if b, ok := input.Input.Input.(bool); ok {
		if b {
			return infer.FunctionResponse[TostringResult]{Output: TostringResult{toStrPtr("true")}}, nil
		}
		return infer.FunctionResponse[TostringResult]{Output: TostringResult{toStrPtr("false")}}, nil
	} else if i, ok := input.Input.Input.(int); ok {
		return infer.FunctionResponse[TostringResult]{Output: TostringResult{toStrPtr(fmt.Sprintf("%v", i))}}, nil
	} else if f, ok := input.Input.Input.(float64); ok {
		return infer.FunctionResponse[TostringResult]{Output: TostringResult{toStrPtr(fmt.Sprintf("%v", f))}}, nil
	}

	return infer.FunctionResponse[TostringResult]{Output: TostringResult{nil}}, fmt.Errorf("%v is not a string value", input.Input.Input)
}

func toStrPtr(s string) *string {
	return &s
}
