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
	"github.com/apparentlymart/go-textseg/textseg"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Strrev struct{}
type StrrevArgs struct {
	Input string `pulumi:"input"`
}

type StrrevResult struct {
	Result string `pulumi:"result"`
}

func (r *Strrev) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the given string with all of its Unicode characters in reverse order.")
}

func (*Strrev) Call(ctx p.Context, args StrrevArgs) (StrrevResult, error) {
	in := []byte(args.Input)
	out := make([]byte, len(in))
	pos := len(out)

	inB := []byte(in)
	for i := 0; i < len(in); {
		d, _, _ := textseg.ScanGraphemeClusters(inB[i:], true)
		cluster := in[i : i+d]
		pos -= len(cluster)
		copy(out[pos:], cluster)
		i += d
	}

	return StrrevResult{string(out)}, nil
}
