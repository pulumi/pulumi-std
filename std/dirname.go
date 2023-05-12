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
	"path/filepath"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Dirname struct{}
type DirnameArgs struct {
	Input string `pulumi:"input"`
}

type DirnameResult struct {
	Result string `pulumi:"result"`
}

func (r *Dirname) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns all but the last element of path, typically the path's directory.")
}

func (*Dirname) Call(_ p.Context, args DirnameArgs) (DirnameResult, error) {
	return DirnameResult{filepath.Dir(args.Input)}, nil
}
