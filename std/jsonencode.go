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

type Jsonencode struct{}
type JsonencodeArgs struct {
	Input interface{} `pulumi:"input"`
}

type JsonencodeResult struct {
	Result string `pulumi:"result"`
}

func (r *Jsonencode) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns a JSON-encoded representation of the given value,
which can contain arbitrarily-nested lists and maps.
Note that if the value is a string then its value will be placed in quotes.`)
}

func (*Jsonencode) Call(_ p.Context, args JsonencodeArgs) (JsonencodeResult, error) {
	inputBytes, err := json.Marshal(args.Input)
	if err != nil {
		return JsonencodeResult{}, err
	}
	return JsonencodeResult{string(inputBytes)}, nil
}
