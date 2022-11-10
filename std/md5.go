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
	//nolint // we need md5 support for compatibility
	"crypto/md5"
	"encoding/hex"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Md5 struct{}
type Md5Args struct {
	Input string `pulumi:"input"`
}

type Md5Result struct {
	Result string `pulumi:"result"`
}

func (r *Md5) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.")
}

var md5AsHex = stringHashFunction(md5.New, hex.EncodeToString)

func (*Md5) Call(ctx p.Context, args Md5Args) (Md5Result, error) {
	return Md5Result{md5AsHex(args.Input)}, nil
}
