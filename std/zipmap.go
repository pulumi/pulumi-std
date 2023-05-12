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

type Zipmap struct{}
type ZipmapArgs struct {
	Keys   []string      `pulumi:"keys"`
	Values []interface{} `pulumi:"values"`
}

type ZipmapResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func (r *Zipmap) Annotate(a infer.Annotator) {
	a.Describe(r, `Constructs a map from a list of keys and a corresponding list of values.`)
}

func (*Zipmap) Call(_ p.Context, args ZipmapArgs) (ZipmapResult, error) {
	if len(args.Keys) != len(args.Values) {
		return ZipmapResult{}, errors.New("keys and values must be the same length")
	}

	result := map[string]interface{}{}
	for i, key := range args.Keys {
		result[key] = args.Values[i]
	}

	return ZipmapResult{result}, nil
}
