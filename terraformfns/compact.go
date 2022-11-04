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

package terraformfns

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Compact struct{}
type CompactArgs struct {
	Input []string `pulumi:"input"`
}

type CompactResult struct {
	Result []string `pulumi:"result"`
}

func (r *Compact) Annotate(a infer.Annotator) {
	a.Describe(r, "Removes empty string elements from a list.")
}

func (*Compact) Call(ctx p.Context, args CompactArgs) (CompactResult, error) {
	output := make([]string, 0)
	for _, value := range args.Input {
		if value != "" {
			output = append(output, value)
		}
	}

	return CompactResult{output}, nil
}
