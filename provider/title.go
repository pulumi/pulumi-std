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

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Title struct{}
type TitleArgs struct {
	Input string `pulumi:"input"`
}

type TitleResult struct {
	Result string `pulumi:"result"`
}

func (r *Title) Annotate(a infer.Annotator) {
	a.Describe(r, "Converts the first letter of each word in the given string to uppercase.")
}

func (*Title) Invoke(
	_ context.Context,
	input infer.FunctionRequest[TitleArgs],
) (infer.FunctionResponse[TitleResult], error) {
	return infer.FunctionResponse[TitleResult]{
		Output: TitleResult{cases.Title(language.English).String(input.Input.Input)},
	}, nil
}
