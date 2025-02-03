// Copyright 2022-2025, Pulumi Corporation.
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
	"math"
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
		"Returns the first non-nil or non-empty value from the given arguments. "+
			"All arguments must be of the same type, or convertible to a common type.",
	)
}

func (*Coalesce) Call(_ p.Context, args CoalesceArgs) (CoalesceResult, error) {
	resultType, err := assignableType(args.Input)
	if err != nil {
		return CoalesceResult{}, err
	}

	if resultType == reflect.TypeOf("") {
		for _, v := range args.Input {
			if v == nil {
				continue
			}

			var s string
			if stringer, ok := v.(fmt.Stringer); ok {
				s = stringer.String()
			} else {
				s = fmt.Sprintf("%v", v)
			}

			if s != "" {
				return CoalesceResult{s}, nil
			}
		}
	} else {
		for _, v := range args.Input {
			if v == nil {
				continue
			}

			result := reflect.ValueOf(v).Convert(resultType).Interface()
			return CoalesceResult{result}, nil
		}
	}

	return CoalesceResult{}, errors.New("no non-empty values found")
}

// assignableType returns a type to which all its arguments are assignable, or an error if one cannot be found.
// assignableType implements Terraform's rules for type conversion as described at
// https://developer.hashicorp.com/terraform/language/expressions/types#type-conversion, augmented to handle the fact
// that Pulumi supports integer types and not just floating-point numbers. Specifically:
//
// * Integers that fit within 32 bits are treated as 32-bit
// * Integers that do not fit within 32 bits are treated as 64-bit
// * All floats are 64-bit
// * Integers that fit with 32 bits can be upcast into 64-bit floats
// * All values can be upcast into strings
func assignableType(args []interface{}) (reflect.Type, error) {
	sawBool := false
	sawInt32 := false
	sawInt64 := false
	sawFloat64 := false
	sawString := false

	for _, arg := range args {
		if arg == nil {
			continue
		}

		switch arg.(type) {
		case bool:
			sawBool = true
			continue
		case int:
			argInt := arg.(int)
			if argInt >= math.MinInt32 && argInt <= math.MaxInt32 {
				sawInt32 = true
			} else {
				sawInt64 = true
			}
			continue
		case int8, int16, int32:
			sawInt32 = true
			continue
		case int64:
			sawInt64 = true
			continue
		case float32, float64:
			sawFloat64 = true
			continue
		case string:
			sawString = true
			continue
		}

		if _, ok := arg.(fmt.Stringer); ok {
			sawString = true
			continue
		}

		return nil, errors.New("all arguments must be of a supported type")
	}

	// If any of the arguments was a string, we can convert all of them to strings.
	if sawString {
		return reflect.TypeOf(""), nil
	}

	if sawBool {
		// If we saw no numbers and we saw at least one boolean, we can convert all of them to booleans.
		if !sawInt64 && !sawFloat64 && !sawInt32 {
			return reflect.TypeOf(true), nil
		}
	} else {
		// If we saw no booleans and we saw at least one number, we can convert to a number type, unless we saw both 64-bit
		// integers and floating-point numbers.
		if sawInt64 && sawFloat64 {
			return nil, errors.New("all arguments must be convertible to the same type")
		}

		if sawInt64 {
			return reflect.TypeOf(int64(0)), nil
		}

		if sawFloat64 {
			return reflect.TypeOf(float64(0.0)), nil
		}

		if sawInt32 {
			return reflect.TypeOf(int32(0)), nil
		}

		// If we saw nothing, return nothing.
		return nil, nil
	}

	// In all other cases, we can't convert.
	return nil, errors.New("all arguments must be convertible to the same type")
}
