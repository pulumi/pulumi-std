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

type Tostring struct{}
type TostringArgs struct {
	Input interface{} `pulumi:"input"`
}

type TostringResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Tostring) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a string value. Only primitive types (string, number, bool)
	and null can be converted to string. All other values will result in an error.`)
}

func (*Tostring) Call(ctx p.Context, args TostringArgs) (TostringResult, error) {
	if args.Input == nil {
		return TostringResult{nil}, nil
	}
	if v, ok := args.Input.(string); ok {
		return TostringResult{v}, nil
	} else if b, ok := args.Input.(bool); ok {
		if b {
			return TostringResult{"true"}, nil
		}
		return TostringResult{"false"}, nil
	} else if i, ok := args.Input.(int); ok {
		return TostringResult{fmt.Sprintf("%v", i)}, nil
	} else if f, ok := args.Input.(float64); ok {
		return TostringResult{fmt.Sprintf("%v", f)}, nil
	}

	return TostringResult{nil}, fmt.Errorf("%v is not a string value", args.Input)
}
