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
	"errors"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Chunklist struct{}
type ChunklistArgs struct {
	Input []interface{} `pulumi:"input"`
	Size  int           `pulumi:"size"`
}

type ChunklistResult struct {
	Result []interface{} `pulumi:"result"`
}

func (r *Chunklist) Annotate(a infer.Annotator) {
	a.Describe(r, "Splits a single list into multiple lists where each has at most the given number of elements.")
}

func (*Chunklist) Call(ctx p.Context, args ChunklistArgs) (ChunklistResult, error) {
	if args.Size < 0 {
		return ChunklistResult{}, errors.New("size must be greater than or equal to 0")
	}

	output := make([]interface{}, 0)
	for i := 0; i < len(args.Input); i += args.Size {
		end := i + args.Size
		if end > len(args.Input) {
			end = len(args.Input)
		}
		output = append(output, args.Input[i:end])
	}

	return ChunklistResult{output}, nil
}
