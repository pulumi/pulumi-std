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
	"strconv"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Tonumber struct{}
type TonumberArgs struct {
	Input interface{} `pulumi:"input"`
}

type TonumberResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Tonumber) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a number value. Only number values, null, and strings
	containing decimal representations of numbers can be converted to number. All other values will result in an error`)
}

func (*Tonumber) Call(ctx p.Context, args TonumberArgs) (TonumberResult, error) {
	if args.Input == nil {
		return TonumberResult{nil}, nil
	}
	if v, ok := args.Input.(int); ok {
		return TonumberResult{v}, nil
	} else if v, ok := args.Input.(float64); ok {
		return TonumberResult{v}, nil
	} else if str, ok := args.Input.(string); ok {
		vInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			vFloat, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return TonumberResult{nil}, fmt.Errorf("%v is not a number value", args.Input)
			}
			return TonumberResult{vFloat}, nil
		}
		return TonumberResult{vInt}, nil
	}
	return TonumberResult{nil}, fmt.Errorf("%v is not a number value", args.Input)
}
