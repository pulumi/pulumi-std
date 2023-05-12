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

type Alltrue struct{}
type AlltrueArgs struct {
	Input []interface{} `pulumi:"input"`
}

type AlltrueResult struct {
	Result bool `pulumi:"result"`
}

func (r *Alltrue) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns true if all elements in a given collection are true or \"true\".
It also returns true if the collection is empty.`)
}

func (*Alltrue) Call(_ p.Context, args AlltrueArgs) (AlltrueResult, error) {
	for _, v := range args.Input {
		value, isBool := v.(bool)
		if isBool && !value {
			return AlltrueResult{false}, nil
		}

		text, isText := v.(string)

		if isText && text != "true" {
			return AlltrueResult{false}, nil
		}

		if !isBool && !isText {
			return AlltrueResult{false}, nil
		}
	}

	return AlltrueResult{true}, nil
}
