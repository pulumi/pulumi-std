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
	//nolint // we need sha1 support for compatibility
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Sha1 struct{}
type Sha1Args struct {
	Input string `pulumi:"input"`
}

type Sha1Result struct {
	Result string `pulumi:"result"`
}

func (r *Sha1) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a hexadecimal representation of the SHA-1 hash of the given string.")
}

var sha1AsHex = stringHashFunction(sha1.New, hex.EncodeToString)

func (*Sha1) Call(_ context.Context, args Sha1Args) (Sha1Result, error) {
	return Sha1Result{sha1AsHex(args.Input)}, nil
}
