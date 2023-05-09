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
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Filesha512 struct{}
type Filesha512Args struct {
	Input string `pulumi:"input"`
}

type Filesha512Result struct {
	Result string `pulumi:"result"`
}

func (r *Filesha512) Annotate(a infer.Annotator) {
	a.Describe(r, "Reads the contents of a file into a string and returns the SHA512 hash of it.")
}

func (*Filesha512) Call(_ p.Context, args Filesha512Args) (Filesha512Result, error) {
	contents, err := readFileContents(args.Input)
	if err != nil {
		return Filesha512Result{}, err
	}
	return Filesha512Result{sha512AsHex(contents)}, nil
}
