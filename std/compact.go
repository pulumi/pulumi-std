// Copyright 2022-2025, Pulumi Corporation.
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

type (
	Compact     struct{}
	CompactArgs struct {
		Input []interface{} `pulumi:"input"`
	}
)

type CompactResult struct {
	Result []string `pulumi:"result"`
}

func (r *Compact) Annotate(a infer.Annotator) {
	a.Describe(r, "Removes empty and nil string elements from a list.")
}

func (*Compact) Call(_ p.Context, args CompactArgs) (CompactResult, error) {
	output := make([]string, 0)
	for _, value := range args.Input {
		if value == nil {
			continue
		}

		v, ok := value.(string)
		if !ok {
			return CompactResult{nil}, errors.New("compact arg is not a string value")
		}

		if ok && v != "" {
			output = append(output, v)
		}
	}

	return CompactResult{output}, nil
}
