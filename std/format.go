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
	"encoding/json"
	"fmt"
	"math"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	Format     struct{}
	FormatArgs struct {
		Input string        `pulumi:"input"`
		Args  []interface{} `pulumi:"args"`
	}
)

type FormatResult struct {
	Result string `pulumi:"result"`
}

func (r *Format) Annotate(a infer.Annotator) {
	a.Describe(r, "Formats a string according to the given format. The syntax for the format is standard sprintf syntax.")
}

// Call implements the `format` function. It is designed to match the behaviour of the Terraform `format` function. The
// implementation of Terraform's function is taken from `go-cty`, and lives at
// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/format.go. When making changes here to ensure
// compatibility, this source can be used as a reference (and is, for instance, the source of existing checks such as
// whether numbers are floating-point or integers).
func (*Format) Call(_ context.Context, args FormatArgs) (FormatResult, error) {
	return FormatResult{format(args.Input, args.Args...)}, nil
}

// format is a helper function that implements the shared logic of the `format` and `formatlist` functions. It takes a
// format string and a list of arguments, and returns the formatted string.
func format(input string, args ...interface{}) string {
	processedArgs := make([]interface{}, len(args))
	for i, arg := range args {
		processedArgs[i] = formattable{value: arg}
	}

	return fmt.Sprintf(input, processedArgs...)
}

// formattable is a type that wraps a value and implements the fmt.Formatter interface. This allows us to control how
// values are formatted when passed to fmt.Sprintf, so that we can match the behaviour of Terraform's `format` function.
type formattable struct {
	value interface{}
}

// Format implements the fmt.Formatter interface's Format function.
func (f formattable) Format(s fmt.State, verb rune) {
	// All numbers come in over the wire as float64, but for integer-based format specifiers such as %b, %o, %x, etc. to
	// work correctly, we'll need to spot these and convert them to integer representations before delegating to
	// fmt.Sprintf. This means that things like %d, will work on values like 1.0, which can be converted into integers.
	// This is the same behaviour as Terraform's format function, though it does differ from e.g. Go's fmt.Sprintf, which
	// will refuse to format a float as an integer.
	v := f.value
	if fv, ok := v.(float64); ok {
		if fv == math.Trunc(fv) && fv <= math.MaxInt64 && fv >= math.MinInt64 {
			v = int64(fv)
		}
	}

	// If we've been asked to print a map or slice, we always convert it to JSON. Technically, we need to convert all
	// values to JSON when the format specifier is #v, but it turns out that all types apart from maps and slices already
	// work correctly in this case.
	if verb == 'v' {
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			bytes, err := json.Marshal(v)
			if err != nil {
				fmt.Fprintf(s, "<error: %v>", err)
				return
			}

			fmt.Fprint(s, string(bytes))
			return
		}
	}

	// In all other cases, we can just delegate to fmt.Sprintf as-is, which should work correctly.
	fmt.Fprintf(s, fmt.FormatString(s, verb), v)
}
