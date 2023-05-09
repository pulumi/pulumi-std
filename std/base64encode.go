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
	"encoding/base64"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Base64encode struct{}
type Base64EncodeArgs struct {
	Input string `pulumi:"input"`
}

type Base64EncodeResult struct {
	Result string `pulumi:"result"`
}

func (r *Base64encode) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a base64-encoded representation of the given string.")
}

func base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func (*Base64encode) Call(_ p.Context, args Base64EncodeArgs) (Base64EncodeResult, error) {
	return Base64EncodeResult{base64Encode(args.Input)}, nil
}
