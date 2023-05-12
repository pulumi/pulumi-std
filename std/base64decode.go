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
	"fmt"
	"unicode/utf8"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Base64decode struct{}
type Base64DecodeArgs struct {
	Input string `pulumi:"input"`
}

type Base64DecodeResult struct {
	Result string `pulumi:"result"`
}

func (r *Base64decode) Annotate(a infer.Annotator) {
	a.Describe(r, "Given a base64-encoded string, decodes it and returns the original string.")
}

func base64Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %w", err)
	}
	if !utf8.Valid(decoded) {
		return "", fmt.Errorf("decoded base64 data is not valid UTF-8")
	}

	return string(decoded), nil
}

func (*Base64decode) Call(_ p.Context, args Base64DecodeArgs) (Base64DecodeResult, error) {
	decoded, err := base64Decode(args.Input)
	if err != nil {
		return Base64DecodeResult{}, err
	}

	return Base64DecodeResult{decoded}, nil
}
