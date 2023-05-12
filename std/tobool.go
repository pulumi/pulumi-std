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

type Tobool struct{}
type ToboolArgs struct {
	Input interface{} `pulumi:"input"`
}

type ToboolResult struct {
	Result *bool `pulumi:"result,optional"`
}

func (r *Tobool) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts its argument to a boolean value. Only boolean values, null, and the exact strings
	"true" and "false" can be converted to boolean. All other values will result in an error.`)
}

func (*Tobool) Call(_ p.Context, args ToboolArgs) (ToboolResult, error) {
	if args.Input == nil {
		return ToboolResult{nil}, nil
	}
	if v, ok := args.Input.(bool); ok {
		return ToboolResult{toBoolPtr(v)}, nil
	} else if str, ok := args.Input.(string); ok {
		v, err := strconv.ParseBool(str)
		if err != nil {
			return ToboolResult{nil}, fmt.Errorf("%v is not a boolean value", args.Input)
		}
		return ToboolResult{toBoolPtr(v)}, nil
	}

	return ToboolResult{nil}, fmt.Errorf("%v is not a boolean value", args.Input)
}

func toBoolPtr(b bool) *bool {
	return &b
}
