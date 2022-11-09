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
	"encoding/hex"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sha512 struct{}
type Sha512Args struct {
	Input string `pulumi:"input"`
}

type Sha512Result struct {
	Result string `pulumi:"result"`
}

func (r *Sha512) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a hexadecimal representation of the SHA-512 hash of the given string.")
}

var sha512AsHex = stringHashFunction(sha512.New, hex.EncodeToString)

func (*Sha512) Call(ctx p.Context, args Sha512Args) (Sha512Result, error) {
	return Sha512Result{sha512AsHex(args.Input)}, nil
}
