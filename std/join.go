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

type Join struct{}
type JoinArgs struct {
	Input     []string `pulumi:"input"`
	Separator string   `pulumi:"separator"`
}

type JoinResult struct {
	Result string `pulumi:"result"`
}

func (r *Join) Annotate(a infer.Annotator) {
	a.Describe(r, "Joins the list with the delimiter for a resultant string.")
}

func (*Join) Invoke(
	_ context.Context,
	input infer.FunctionRequest[JoinArgs],
) (infer.FunctionResponse[JoinResult], error) {
	return infer.FunctionResponse[JoinResult]{
		Output: JoinResult{strings.Join(input.Input.Input, input.Input.Separator)},
	}, nil
}
