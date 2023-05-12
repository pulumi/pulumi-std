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
	"encoding/json"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Jsondecode struct{}
type JsondecodeArgs struct {
	Input string `pulumi:"input"`
}

type JsondecodeResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Jsondecode) Annotate(a infer.Annotator) {
	a.Describe(r, `Interprets a given string as JSON and returns a represetation
	of the result of decoding that string.
	If input is not valid JSON, the result will be the input unchanged.`)
}

func (*Jsondecode) Call(_ p.Context, args JsondecodeArgs) (JsondecodeResult, error) {
	var res map[string]interface{}
	if err := json.Unmarshal([]byte(args.Input), &res); err != nil {
		return JsondecodeResult{args.Input}, nil

	}
	return JsondecodeResult{res}, nil
}
