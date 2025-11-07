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
	"context"
	"errors"
	"fmt"
	"reflect"

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

func (*Coalesce) Call(_ context.Context, args CoalesceArgs) (CoalesceResult, error) {
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
	} else if resultType != nil {
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
