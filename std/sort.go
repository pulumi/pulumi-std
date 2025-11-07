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
	"sort"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sort struct{}
type SortArgs struct {
	Input []string `pulumi:"input"`
}

type SortResult struct {
	Result []string `pulumi:"result"`
}

func (r *Sort) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns a list of strings sorted lexicographically.`)
}

func (*Sort) Invoke(_ context.Context, input infer.FunctionRequest[SortArgs]) (infer.FunctionResponse[SortResult], error) {
	sort.Strings(input.Input.Input)
	return infer.FunctionResponse[SortResult]{Output: SortResult{input.Input.Input}}, nil
}
