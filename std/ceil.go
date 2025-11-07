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
	"math"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Ceil struct{}
type CeilArgs struct {
	Input float64 `pulumi:"input"`
}

type CeilResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Ceil) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the least integer value greater than or equal to the argument.")
}

func (*Ceil) Invoke(_ context.Context, input infer.FunctionRequest[CeilArgs]) (infer.FunctionResponse[CeilResult], error) {
	return infer.FunctionResponse[CeilResult]{Output: CeilResult{math.Ceil(input.Input.Input)}}, nil
}
