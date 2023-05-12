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

type Abspath struct{}
type AbspathArgs struct {
	Input string `pulumi:"input"`
}

type AbspathResult struct {
	Result string `pulumi:"result"`
}

func (r *Abspath) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns an absolute representation of the specified path.
If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.`)
}

func (*Abspath) Call(_ p.Context, args AbspathArgs) (AbspathResult, error) {
	abs, err := filepath.Abs(args.Input)
	return AbspathResult{abs}, err
}
