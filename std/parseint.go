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

type Parseint struct{}
type ParseintArgs struct {
	Input string `pulumi:"input"`
	Base  *int   `pulumi:"base,optional"`
}

type ParseintResult struct {
	Result int `pulumi:"result"`
}

func (r *Parseint) Annotate(a infer.Annotator) {
	a.Describe(r, `Parses the given string as a representation of an integer in the specified base
and returns the resulting number. The base must be between 2 and 62 inclusive.
	.`)
}

func (*Parseint) Invoke(_ context.Context, input infer.FunctionRequest[ParseintArgs]) (infer.FunctionResponse[ParseintResult], error) {
	base := 10
	if input.Input.Base != nil {
		base = *input.Input.Base
	}
	if base < 2 || base > 62 {
		return infer.FunctionResponse[ParseintResult]{Output: ParseintResult{}}, fmt.Errorf("base must be between 2 and 62 inclusive")
	}
	result, err := strconv.ParseInt(input.Input.Input, base, 64)
	if err != nil {
		return infer.FunctionResponse[ParseintResult]{Output: ParseintResult{}}, err
	}
	return infer.FunctionResponse[ParseintResult]{Output: ParseintResult{int(result)}}, nil
}
