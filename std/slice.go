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
	"errors"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Slice struct{}
type SliceArgs struct {
	List []interface{} `pulumi:"list"`
	From *int          `pulumi:"from,optional"`
	To   *int          `pulumi:"to,optional"`
}

type SliceResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Slice) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the portion of list between from (inclusive) and to (exclusive).")
}

func (*Slice) Invoke(_ context.Context, input infer.FunctionRequest[SliceArgs]) (infer.FunctionResponse[SliceResult], error) {
	from := 0
	if input.Input.From != nil {
		from = *input.Input.From
	}

	to := len(input.Input.List)
	if input.Input.To != nil {
		to = *input.Input.To
	}

	if from < 0 {
		from += len(input.Input.List)
	}
	if to < 0 {
		to += len(input.Input.List)
	}

	if from < 0 || from > len(input.Input.List) {
		return infer.FunctionResponse[SliceResult]{Output: SliceResult{}}, errors.New("from index out of range")
	}
	if to < 0 || to > len(input.Input.List) {
		return infer.FunctionResponse[SliceResult]{Output: SliceResult{}}, errors.New("to index out of range")
	}

	return infer.FunctionResponse[SliceResult]{Output: SliceResult{input.Input.List[from:to]}}, nil
}
