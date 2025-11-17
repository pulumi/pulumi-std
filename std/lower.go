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
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Lower struct{}
type LowerArgs struct {
	Input string `pulumi:"input"`
}

type LowerResult struct {
	Result string `pulumi:"result"`
}

func (r *Lower) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a copy of the string with all Unicode letters mapped to their lower case.")
}

func (*Lower) Invoke(
	_ context.Context,
	input infer.FunctionRequest[LowerArgs],
) (infer.FunctionResponse[LowerResult], error) {
	return infer.FunctionResponse[LowerResult]{Output: LowerResult{strings.ToLower(input.Input.Input)}}, nil
}
