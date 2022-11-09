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

type Substr struct{}
type SubstrArgs struct {
	Input  string `pulumi:"input"`
	Offset int    `pulumi:"offset"`
	Length int    `pulumi:"length"`
}

type SubstrResult struct {
	Result string `pulumi:"result"`
}

func (r *Substr) Annotate(a infer.Annotator) {
	a.Describe(r, "Extracts a substring from the given string.")
}

func (*Substr) Call(ctx p.Context, args SubstrArgs) (SubstrResult, error) {
	in := []byte(args.Input)
	if args.Length == 0 {
		return SubstrResult{""}, nil
	}

	for args.Offset < 0 {
		args.Offset += len(args.Input)
	}

	sub := in
	pos := 0
	var i int

	if args.Offset > 0 {
		for i = 0; i < len(sub); {
			d, _, _ := textseg.ScanGraphemeClusters(sub[i:], true)
			i += d
			pos++
			if pos == args.Offset {
				break
			}
			if i >= len(in) {
				return SubstrResult{""}, nil
			}
		}

		sub = sub[i:]
	}

	if args.Length < 0 {
		return SubstrResult{string(sub)}, nil
	}

	pos = 0
	for i = 0; i < len(sub); {
		d, _, _ := textseg.ScanGraphemeClusters(sub[i:], true)
		i += d
		pos++
		if pos == args.Length {
			break
		}
	}

	sub = sub[:i]
	return SubstrResult{string(sub)}, nil
}
