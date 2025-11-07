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
	"fmt"

	"github.com/pulumi/pulumi-go-provider/infer"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}
type BcryptArgs struct {
	Input string `pulumi:"input"`
	Cost  *int   `pulumi:"cost,optional"`
}

type BcryptResult struct {
	Result string `pulumi:"result"`
}

func (r *Bcrypt) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns the Blowfish encrypted hash of the string at the given cost.
A default cost of 10 will be used if not provided.`)
}

func (*Bcrypt) Invoke(_ context.Context, input infer.FunctionRequest[BcryptArgs]) (infer.FunctionResponse[BcryptResult], error) {
	defaultCost := 10
	if input.Input.Cost != nil {
		defaultCost = *input.Input.Cost
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Input.Input), defaultCost)
	if err != nil {
		return infer.FunctionResponse[BcryptResult]{Output: BcryptResult{}}, fmt.Errorf("error occurred generating password %s", err.Error())
	}
	return infer.FunctionResponse[BcryptResult]{Output: BcryptResult{string(hash)}}, nil
}
