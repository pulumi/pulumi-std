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
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sum struct{}
type SumArgs struct {
	Input []float64 `pulumi:"input"`
}

type SumResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Sum) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the total sum of the elements of the input list.")
}

func (*Sum) Call(_ p.Context, args SumArgs) (SumResult, error) {
	sum := 0.0
	for _, current := range args.Input {
		sum += current
	}

	return SumResult{sum}, nil
}
