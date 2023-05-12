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
	"regexp"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Chomp struct{}
type ChompArgs struct {
	Input string `pulumi:"input"`
}

type ChompResult struct {
	Result string `pulumi:"result"`
}

func (r *Chomp) Annotate(a infer.Annotator) {
	a.Describe(r, "Removes one or more newline characters from the end of the given string.")
}

func (*Chomp) Call(_ p.Context, args ChompArgs) (ChompResult, error) {
	newlines := regexp.MustCompile(`(?:\r\n?|\n)*\z`)
	result := newlines.ReplaceAllString(args.Input, "")
	return ChompResult{result}, nil
}
