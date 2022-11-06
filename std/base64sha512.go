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
	"crypto/sha512"
	"encoding/base64"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Base64sha512 struct{}
type Base64Sha512Args struct {
	Input string `pulumi:"input"`
}

type Base64Sha512Result struct {
	Result string `pulumi:"result"`
}

func (r *Base64sha512) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns a base64-encoded representation of raw SHA-512 sum of the given string. 
This is not equivalent of base64encode(sha512(string)) since sha512() returns hexadecimal representation.`)
}

var base64Sha512 = stringHashFunction(sha512.New, base64.StdEncoding.EncodeToString)

func (*Base64sha512) Call(ctx p.Context, args Base64Sha512Args) (Base64Sha512Result, error) {
	return Base64Sha512Result{base64Sha512(args.Input)}, nil
}
