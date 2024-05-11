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
	"time"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Unixtorfc3999 struct{}
type Unixtorfc3999args struct {
	Input string `pulumi:"input"`
}

type Unixtorfc3999result struct {
	Result int64 `pulumi:"result"`
}

func (r *Unixtorfc3999) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.`)
}

func (*Unixtorfc3999) Call(_ p.Context, args Unixtorfc3999args) (Unixtorfc3999result, error) {

	t, err := time.Parse(time.RFC3339, args.Input)
	if err != nil {
		return Unixtorfc3999result{}, err
	}
	return Unixtorfc3999result{t.UnixMilli()}, nil
}
