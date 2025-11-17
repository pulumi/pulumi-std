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
	"reflect"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	Setintersection     struct{}
	SetintersectionArgs struct {
		Input [][]interface{} `pulumi:"input"`
	}
)

type SetintersectionResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Setintersection) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the set of elements that all the input sets have in common.")
}

// Call implements the `setintersection` function. It is designed to match the behaviour of the Terraform
// `setintersection` function. That said, it currently does not support values of sets which are lists or maps/objects.
// In practice this is hopefully not a massive issue -- Terraform itself is not super consistent in how it treats
// compound literals like this (e.g. `setintersection([[1, 2]], [[2, 1]])` returns `[[]]` -- the inner lists are *not*
// treated as sets, while the outer ones are).
func (*Setintersection) Invoke(
	_ context.Context,
	input infer.FunctionRequest[SetintersectionArgs],
) (infer.FunctionResponse[SetintersectionResult], error) {
	if len(input.Input.Input) == 0 {
		return infer.FunctionResponse[SetintersectionResult]{
			Output: SetintersectionResult{Result: []interface{}{}},
		}, nil
	}

	allInputs := []interface{}{}
	for _, set := range input.Input.Input {
		allInputs = append(allInputs, set...)
	}

	resultType, err := assignableType(allInputs)
	if resultType == nil || err != nil {
		return infer.FunctionResponse[SetintersectionResult]{
			Output: SetintersectionResult{Result: []interface{}{}},
		}, err
	}

	seen := map[interface{}]int{}
	resultTypeIsString := resultType == reflect.TypeOf("")

	for i, set := range input.Input.Input {
		for _, item := range set {
			if resultTypeIsString {
				if stringer, ok := item.(fmt.Stringer); ok {
					item = stringer.String()
				} else {
					item = fmt.Sprintf("%v", item)
				}
			} else {
				item = reflect.ValueOf(item).Convert(resultType).Interface()
			}

			count := seen[item]
			if count == i {
				seen[item] = i + 1
			}
		}
	}

	result := []interface{}{}
	requiredCount := len(input.Input.Input)
	for item, count := range seen {
		if count == requiredCount {
			result = append(result, item)
		}
	}

	return infer.FunctionResponse[SetintersectionResult]{Output: SetintersectionResult{Result: result}}, nil
}
