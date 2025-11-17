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
	"crypto/sha256"
	"encoding/hex"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sha256 struct{}
type Sha256Args struct {
	Input string `pulumi:"input"`
}

type Sha256Result struct {
	Result string `pulumi:"result"`
}

func (r *Sha256) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a hexadecimal representation of the SHA-256 hash of the given string.")
}

var sha256AsHex = stringHashFunction(sha256.New, hex.EncodeToString)

func (*Sha256) Invoke(
	_ context.Context,
	input infer.FunctionRequest[Sha256Args],
) (infer.FunctionResponse[Sha256Result], error) {
	return infer.FunctionResponse[Sha256Result]{Output: Sha256Result{sha256AsHex(input.Input.Input)}}, nil
}
