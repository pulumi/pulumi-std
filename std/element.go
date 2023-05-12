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

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Element struct{}
type ElementArgs struct {
	Input []interface{} `pulumi:"input"`
	Index int           `pulumi:"index"`
}

type ElementResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Element) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the element at the specified index.")
}

func (*Element) Call(_ p.Context, args ElementArgs) (ElementResult, error) {
	if len(args.Input) == 0 {
		return ElementResult{}, errors.New("input list must not be empty when using the element function")
	}

	index := args.Index % len(args.Input)
	for index < 0 {
		index += len(args.Input)
	}

	return ElementResult{args.Input[index]}, nil
}
