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

type Merge struct{}
type MergeArgs struct {
	Input []map[string]interface{} `pulumi:"input"`
}

type MergeResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func (r *Merge) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns the union of 2 or more maps. The maps are consumed in the order provided,
and duplicate keys overwrite previous entries.`)
}

func (*Merge) Call(_ context.Context, args MergeArgs) (MergeResult, error) {
	result := make(map[string]interface{})
	for _, m := range args.Input {
		for k, v := range m {
			result[k] = v
		}
	}
	return MergeResult{result}, nil
}
