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
	"errors"
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Map struct{}
type MapArgs struct {
	Args []interface{} `pulumi:"args"`
}

type MapResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func (r *Map) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a map consisting of the key/value pairs specified as arguments.")
}

func (*Map) Call(_ context.Context, args MapArgs) (MapResult, error) {
	if len(args.Args)%2 != 0 {
		return MapResult{}, errors.New("expected an even number of arguments")
	}

	result := make(map[string]interface{})
	for i := 0; i < len(args.Args); i += 2 {
		key, ok := args.Args[i].(string)
		if !ok {
			return MapResult{}, fmt.Errorf("expected a string key at index %d", i)
		}
		result[key] = args.Args[i+1]
	}

	return MapResult{result}, nil
}
