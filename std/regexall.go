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
	"errors"
	"regexp"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	Regexall     struct{}
	RegexallArgs struct {
		Pattern string `pulumi:"pattern"`
		String  string `pulumi:"string"`
	}
)

type RegexallResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Regexall) Annotate(a infer.Annotator) {
	a.Describe(
		r,
		"Returns a list of all matches of a regular expression in a string (including named or indexed groups).",
	)
}

func (*Regexall) Invoke(
	_ context.Context,
	input infer.FunctionRequest[RegexallArgs],
) (infer.FunctionResponse[RegexallResult], error) {
	re, err := regexp.Compile(input.Input.Pattern)
	if err != nil {
		return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{}}, err
	}
	matches := re.FindAllStringSubmatch(input.Input.String, -1)

	if matches == nil {
		return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{Result: []interface{}{}}}, nil
	}

	result := make([]interface{}, len(matches))
	// If there are no groups/submatches, return the full match
	if re.NumSubexp() == 0 {
		for i, match := range matches {
			result[i] = match[0]
		}
		return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{Result: result}}, nil
	}

	hasNamedSubExp := false
	for _, name := range re.SubexpNames() {
		if name == "" && hasNamedSubExp {
			return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{}},
				errors.New(
					"regex pattern contains a mix of named and unnamed submatches, must be all named or all unnamed",
				)
		}

		if name != "" {
			hasNamedSubExp = true
		}
	}

	// If there are indexed submatches/capture groups return a list of lists of submatches
	if !hasNamedSubExp {
		for i, match := range matches {
			// do not include the full match, just submatches
			result[i] = interface{}(match[1:])
		}
		return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{Result: result}}, nil
	}

	// If there are named submatches return a map of the named submatches
	for i, match := range matches {
		submatches := make(map[string]string)
		for j, name := range re.SubexpNames() {
			// As per the docs, the first element of the slice is the full match,
			// which has no subgroup so skip it.
			if j != 0 {
				submatches[name] = match[j]
			}
		}

		result[i] = interface{}(submatches)
	}

	return infer.FunctionResponse[RegexallResult]{Output: RegexallResult{Result: result}}, nil
}
