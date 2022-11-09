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
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type TrimSpace struct{}
type TrimspaceArgs struct {
	Input string `pulumi:"input"`
}

type TrimspaceResult struct {
	Result string `pulumi:"result"`
}

func (r *TrimSpace) Annotate(a infer.Annotator) {
	a.Describe(r, `Removes any space characters from the start and end of the given string,
	following the Unicode definition of \"space\" (i.e. spaces, tabs, newline, etc.).`)
}

func (*TrimSpace) Call(ctx p.Context, args TrimspaceArgs) (TrimspaceResult, error) {
	return TrimspaceResult{strings.TrimSpace(args.Input)}, nil
}
