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
	"net/url"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type URLEncode struct{}
type URLencodeArgs struct {
	Input string `pulumi:"input"`
}

type URLencodeResult struct {
	Result string `pulumi:"result"`
}

func (r *URLEncode) Annotate(a infer.Annotator) {
	a.Describe(r, `Applies URL encoding to a given string.`)
}

func (*URLEncode) Call(ctx p.Context, args URLencodeArgs) (URLencodeResult, error) {
	return URLencodeResult{url.QueryEscape(args.Input)}, nil
}
