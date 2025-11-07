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

type Fileexists struct{}
type FileexistsArgs struct {
	Input string `pulumi:"input"`
}

type FileexistsResult struct {
	Result bool `pulumi:"result"`
}

func (r *Fileexists) Annotate(a infer.Annotator) {
	a.Describe(r, "Determines whether a file exists at a given path.")
}

func (*Fileexists) Invoke(_ context.Context, input infer.FunctionRequest[FileexistsArgs]) (infer.FunctionResponse[FileexistsResult], error) {
	_, err := readFileContents(input.Input.Input)
	if err != nil {
		return infer.FunctionResponse[FileexistsResult]{Output: FileexistsResult{false}}, err
	}
	return infer.FunctionResponse[FileexistsResult]{Output: FileexistsResult{true}}, nil
}
