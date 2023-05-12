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
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Anytrue struct{}
type AnytrueArgs struct {
	Input []interface{} `pulumi:"input"`
}

type AnytrueResult struct {
	Result bool `pulumi:"result"`
}

func (r *Anytrue) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns true if any of the elements in a given collection are true or \"true\".
It also returns false if the collection is empty.`)
}

func (*Anytrue) Call(_ p.Context, args AnytrueArgs) (AnytrueResult, error) {
	for _, v := range args.Input {
		value, isBool := v.(bool)
		if isBool && value {
			return AnytrueResult{true}, nil
		}

		text, isText := v.(string)
		if isText && text == "true" {
			return AnytrueResult{true}, nil
		}
	}

	return AnytrueResult{false}, nil
}
