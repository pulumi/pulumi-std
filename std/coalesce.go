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
	"errors"
	"fmt"
	"reflect"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	Coalesce     struct{}
	CoalesceArgs struct {
		Input []interface{} `pulumi:"input"`
	}
)

type CoalesceResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Coalesce) Annotate(a infer.Annotator) {
	a.Describe(
		r,
		"Returns the first non-nil value or non empty string from the given arguments as a the most generic type.",
	)
}

func (*Coalesce) Call(_ p.Context, args CoalesceArgs) (CoalesceResult, error) {
	input := args.Input
	if !allSameType(input) {
		resultType, err := mostGenericType(args.Input)
		if err != nil {
			return CoalesceResult{}, errors.New("all arguments must be of the same type")
		}

		// convert the value to most generic type.
		for i, v := range input {
			if resultType == reflect.TypeOf((*fmt.Stringer)(nil)) {
				input[i] = v.(fmt.Stringer).String()
				continue
			}

			input[i] = reflect.ValueOf(v).Convert(resultType).Interface()
		}
	}

	for _, value := range input {
		if value == nil {
			continue
		}

		switch value := value.(type) {
		case string:
			if value == "" {
				continue
			}
			return CoalesceResult{value}, nil
		default:
			return CoalesceResult{value}, nil
		}
	}

	return CoalesceResult{}, errors.New("no non-empty values found")
}

func allSameType(args []interface{}) bool {
	var t reflect.Type
	for _, arg := range args {
		if arg != nil {
			if t == nil {
				t = reflect.TypeOf(arg)
			} else {
				if reflect.TypeOf(arg) != t {
					return false
				}
			}
		}
	}

	return true
}

func mostGenericType(args []interface{}) (reflect.Type, error) {
	var resultType reflect.Type
	for _, arg := range args {
		if arg == nil {
			continue
		}

		if resultType == nil {
			resultType = reflect.TypeOf(arg)
			continue
		}

		t := reflect.TypeOf(arg)
		if t.AssignableTo(resultType) {
			// if t is a float64 or float32 then make sure mostGenericType is also a float type and down convert to int.
			if t == reflect.TypeOf(float64(0)) || t == reflect.TypeOf(float32(0)) {
				resultType = reflect.TypeOf(float64(0))
			}
			continue
		}

		if resultType.AssignableTo(t) {
			resultType = t
			continue
		}

		// if t implements fmt.Stringer then set mostGenericType to string.
		if t.Implements(reflect.TypeOf((*fmt.Stringer)(nil))) {
			resultType = reflect.TypeOf("")
			continue
		}

		return nil, errors.New("all arguments must be convertable to same type")
	}

	return resultType, nil
}
