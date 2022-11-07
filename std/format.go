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
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Format struct{}
type FormatArgs struct {
	Input string        `pulumi:"input"`
	Args  []interface{} `pulumi:"args"`
}

type FormatResult struct {
	Result string `pulumi:"result"`
}

func (r *Format) Annotate(a infer.Annotator) {
	a.Describe(r, "Formats a string according to the given format. The syntax for the format is standard sprintf syntax.")
}

func (*Format) Call(ctx p.Context, args FormatArgs) (FormatResult, error) {
	return FormatResult{fmt.Sprintf(args.Input, args.Args...)}, nil
}
