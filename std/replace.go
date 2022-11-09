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
	"strings"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Replace struct{}
type ReplaceArgs struct {
	Text    string `pulumi:"text"`
	Search  string `pulumi:"search"`
	Replace string `pulumi:"replace"`
}

type ReplaceResult struct {
	Result string `pulumi:"result"`
}

func (r *Replace) Annotate(a infer.Annotator) {
	a.Describe(r, `Does a search and replace on the given string. 
All instances of search are replaced with the value of replace. 
If search is wrapped in forward slashes, it is treated as a regular expression. 
If using a regular expression, replace can reference subcaptures in the regular expression by 
using $n where n is the index or name of the subcapture. If using a regular expression, 
the syntax conforms to the re2 regular expression syntax.`)
}

func (*Replace) Call(ctx p.Context, args ReplaceArgs) (ReplaceResult, error) {
	if len(args.Search) > 1 && args.Search[0] == '/' && args.Search[len(args.Search)-1] == '/' {
		re, err := regexp.Compile(args.Search[1 : len(args.Search)-1])
		if err != nil {
			return ReplaceResult{}, err
		}

		return ReplaceResult{re.ReplaceAllString(args.Text, args.Replace)}, nil
	}

	return ReplaceResult{strings.Replace(args.Text, args.Search, args.Replace, -1)}, nil
}
