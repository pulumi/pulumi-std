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
	"errors"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Length struct{}
type LengthArgs struct {
	Input interface{} `pulumi:"input"`
}

type LengthResult struct {
	Result int `pulumi:"result"`
}

func (r *Length) Annotate(a infer.Annotator) {
	a.Describe(r, "Determines the length of a given list, map, or string.")
}

func (*Length) Invoke(
	_ context.Context,
	input infer.FunctionRequest[LengthArgs],
) (infer.FunctionResponse[LengthResult], error) {
	switch v := input.Input.Input.(type) {
	case string:
		return infer.FunctionResponse[LengthResult]{Output: LengthResult{len(v)}}, nil
	case []interface{}:
		return infer.FunctionResponse[LengthResult]{Output: LengthResult{len(v)}}, nil
	case map[string]interface{}:
		return infer.FunctionResponse[LengthResult]{Output: LengthResult{len(v)}}, nil
	default:
		return infer.FunctionResponse[LengthResult]{}, errors.New("input must be a string, list, or map")
	}
}
