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
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Split struct{}
type SplitArgs struct {
	Separator string `pulumi:"separator"`
	Text      string `pulumi:"text"`
}

type SplitResult struct {
	Result []string `pulumi:"result"`
}

func (r *Split) Annotate(a infer.Annotator) {
	a.Describe(r, `Produces a list by dividing a given string at all occurrences of a given separator`)
}

func (*Split) Call(_ context.Context, args SplitArgs) (SplitResult, error) {
	return SplitResult{strings.Split(args.Text, args.Separator)}, nil
}
