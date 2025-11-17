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

type Log struct{}
type LogArgs struct {
	Input float64 `pulumi:"input"`
	Base  float64 `pulumi:"base"`
}

type LogResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Log) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the greatest integer value less than or equal to the argument.")
}

func (*Log) Invoke(_ context.Context, input infer.FunctionRequest[LogArgs]) (infer.FunctionResponse[LogResult], error) {
	return infer.FunctionResponse[LogResult]{
		Output: LogResult{math.Log(input.Input.Input) / math.Log(input.Input.Base)},
	}, nil
}
