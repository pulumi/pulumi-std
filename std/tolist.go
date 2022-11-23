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
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Tolist struct{}
type TolistArgs struct {
	Input []interface{} `pulumi:"input"`
}

type TolistResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Tolist) Annotate(a infer.Annotator) {
	a.Describe(r, "Converts its argument to a list value.")
}

func (*Tolist) Call(ctx p.Context, args TolistArgs) (TolistResult, error) {
	if len(args.Input) == 0 {
		return TolistResult{make([]interface{}, 0)}, nil
	}
	return TolistResult{convertListSameType(args.Input)}, nil
}

func isSameType(a, b interface{}) bool {
	return fmt.Sprintf("%T", a) == fmt.Sprintf("%T", b)
}

func toStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func convertListSameType(input []interface{}) []interface{} {
	res := make([]interface{}, len(input))
	sameType := true
	for i := 1; i < len(input); i++ {
		if !isSameType(input, input[i-1]) {
			sameType = false
			break
		}
	}
	if sameType {
		return input
	}
	for i, e := range input {
		res[i] = toStr(e)
	}
	return res
}
