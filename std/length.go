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

type Length struct{}
type LengthArgs struct {
	Input interface{} `pulumi:"input"`
}

type LengthResult struct {
	Result int `pulumi:"result"`
}

func (r *Length) Annotate(a infer.Annotator) {
	a.Describe(r, "Determines the length of a given list, map, or string.")
}

func (*Length) Call(ctx p.Context, args LengthArgs) (LengthResult, error) {
	switch v := args.Input.(type) {
	case string:
		return LengthResult{len(v)}, nil
	case []interface{}:
		return LengthResult{len(v)}, nil
	case map[string]interface{}:
		return LengthResult{len(v)}, nil
	default:
		return LengthResult{}, errors.New("input must be a string, list, or map")
	}
}
