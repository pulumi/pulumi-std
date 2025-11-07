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

	"github.com/apparentlymart/go-textseg/textseg"
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

func (*Strrev) Invoke(_ context.Context, input infer.FunctionRequest[StrrevArgs]) (infer.FunctionResponse[StrrevResult], error) {
	in := []byte(input.Input.Input)
	out := make([]byte, len(in))
	pos := len(out)

	for i := 0; i < len(in); {
		d, _, _ := textseg.ScanGraphemeClusters(in[i:], true)
		cluster := in[i : i+d]
		pos -= len(cluster)
		copy(out[pos:], cluster)
		i += d
	}

	return infer.FunctionResponse[StrrevResult]{Output: StrrevResult{string(out)}}, nil
}
