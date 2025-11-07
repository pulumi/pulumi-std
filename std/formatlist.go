// Copyright 2025, Pulumi Corporation.
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

type (
	Formatlist     struct{}
	FormatlistArgs struct {
		Input string        `pulumi:"input"`
		Args  []interface{} `pulumi:"args"`
	}
)

type FormatlistResult struct {
	Result []string `pulumi:"result"`
}

func (r *Formatlist) Annotate(a infer.Annotator) {
	a.Describe(
		r,
		"Formats a list of strings according to the given format. "+
			"Argument values which are lists are \"zipped\" together to produce a list of results.",
	)
}

func (*Formatlist) Invoke(_ context.Context, input infer.FunctionRequest[FormatlistArgs]) (infer.FunctionResponse[FormatlistResult], error) {
	// If we weren't passed any arguments, we can only attempt to format the input string.
	if len(input.Input.Args) == 0 {
		return infer.FunctionResponse[FormatlistResult]{Output: FormatlistResult{[]string{format(input.Input.Input)}}}, nil
	}

	// We have arguments -- ensure that all lists which appear are of the same length, and calculate that length.
	resultLength := -1
	for _, arg := range input.Input.Args {
		if list, ok := arg.([]interface{}); ok {
			if resultLength == -1 {
				resultLength = len(list)
			} else if len(list) != resultLength {
				return infer.FunctionResponse[FormatlistResult]{Output: FormatlistResult{}}, fmt.Errorf("all list arguments must be of the same length")
			}
		}
	}

	if resultLength == -1 {
		// There are no lists in the arguments, so we can just format the input string.
		return infer.FunctionResponse[FormatlistResult]{Output: FormatlistResult{[]string{format(input.Input.Input, input.Input.Args...)}}}, nil
	}

	if resultLength == 0 {
		// There are lists with no elements in them, so we can return an empty list (this matches the behaviour of
		// Terraform's formatlist).
		return infer.FunctionResponse[FormatlistResult]{Output: FormatlistResult{[]string{}}}, nil
	}

	// We have lists with elements in them, so we need to "zip" them together and format each combination. We do this by
	// looping from 0 to the length we know all the lists share, building a set of arguments for format at each step. The
	// ith set of arguments consists of the ith element of each list, or the original argument if it's not a list.
	results := make([]string, resultLength)
	for i := 0; i < resultLength; i++ {
		callArgs := make([]interface{}, len(input.Input.Args))
		for j, arg := range input.Input.Args {
			if list, ok := arg.([]interface{}); ok {
				callArgs[j] = list[i]
			} else {
				callArgs[j] = arg
			}
		}

		results[i] = format(input.Input.Input, callArgs...)
	}

	return infer.FunctionResponse[FormatlistResult]{Output: FormatlistResult{results}}, nil
}
