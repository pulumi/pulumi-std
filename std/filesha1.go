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
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Filesha1 struct{}
type Filesha1Args struct {
	Input string `pulumi:"input"`
}

type Filesha1Result struct {
	Result string `pulumi:"result"`
}

func (r *Filesha1) Annotate(a infer.Annotator) {
	a.Describe(r, "Reads the contents of a file into a string and returns the SHA1 hash of it.")
}

func (*Filesha1) Invoke(
	_ context.Context,
	input infer.FunctionRequest[Filesha1Args],
) (infer.FunctionResponse[Filesha1Result], error) {
	contents, err := readFileContents(input.Input.Input)
	if err != nil {
		return infer.FunctionResponse[Filesha1Result]{Output: Filesha1Result{}}, err
	}
	return infer.FunctionResponse[Filesha1Result]{Output: Filesha1Result{sha1AsHex(contents)}}, nil
}
