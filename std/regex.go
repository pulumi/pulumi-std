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
	Regex     struct{}
	RegexArgs struct {
		Pattern string `pulumi:"pattern"`
		String  string `pulumi:"string"`
	}
)

type RegexResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Regex) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the first match of a regular expression in a string (including named or indexed groups).")
}

func (*Regex) Call(_ context.Context, args RegexArgs) (RegexResult, error) {
	re, err := regexp.Compile(args.Pattern)
	if err != nil {
		return RegexResult{}, err
	}
	match := re.FindStringSubmatch(args.String)

	if match == nil {
		return RegexResult{Result: []interface{}{}}, nil
	}

	// If there are no groups/submatches, return the full match
	if re.NumSubexp() == 0 {
		return RegexResult{Result: interface{}(match[0])}, nil
	}

	hasNamedSubExp := false
	for _, name := range re.SubexpNames() {
		if name == "" && hasNamedSubExp {
			return RegexResult{},
				errors.New("regex pattern contains a mix of named and unnamed submatches, must be all named or all unnamed")
		}

		if name != "" {
			hasNamedSubExp = true
		}
	}

	// If there are indexed submatches/capture groups return a list of lists of submatches
	if !hasNamedSubExp {
		result := make([]interface{}, 0, len(match)-1)
		for i, submatch := range match {
			if i == 0 {
				// do not include the full match, just submatches
				continue
			}
			result = append(result, submatch)
		}
		return RegexResult{Result: interface{}(result)}, nil
	}

	// If there are named submatches return a map of the named submatches
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i == 0 {
			// do not include the full match, just submatches
			continue
		}
		result[name] = match[i]
	}

	return RegexResult{Result: interface{}(result)}, nil
}
