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
	"math"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Pow struct{}
type PowArgs struct {
	Base     float64 `pulumi:"base"`
	Exponent float64 `pulumi:"exponent"`
}

type PowResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Pow) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the base input raised to the power of the exponent.")
}

func (*Pow) Call(_ p.Context, args PowArgs) (PowResult, error) {
	return PowResult{math.Pow(args.Base, args.Exponent)}, nil
}
