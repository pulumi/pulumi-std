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

type Reverse struct{}
type ReverseArgs struct {
	Input []interface{} `pulumi:"input"`
}

type ReverseResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Reverse) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns a sequence with the same elements but in reverse order.`)
}

func reverse(input []interface{}) {
	l := len(input)
	m := l / 2
	for i := 0; i < m; i++ {
		j := l - i - 1
		input[i], input[j] = input[j], input[i]
	}
}

func (*Reverse) Call(_ context.Context, args ReverseArgs) (ReverseResult, error) {
	reverse(args.Input)
	return ReverseResult{args.Input}, nil
}
