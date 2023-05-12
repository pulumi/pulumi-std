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

type Lookup struct{}
type LookupArgs struct {
	Map     map[string]interface{} `pulumi:"map"`
	Key     string                 `pulumi:"key"`
	Default interface{}            `pulumi:"default,optional"`
}

type LookupResult struct {
	Result interface{} `pulumi:"result"`
}

func (r *Lookup) Annotate(a infer.Annotator) {
	a.Describe(r, "Performs a dynamic lookup into a map variable.")
}

func (*Lookup) Call(_ p.Context, args LookupArgs) (LookupResult, error) {
	for key, value := range args.Map {
		if key == args.Key {
			return LookupResult{value}, nil
		}
	}

	if args.Default != nil {
		return LookupResult{args.Default}, nil
	}

	return LookupResult{nil}, errors.New("key not found in input map")
}
