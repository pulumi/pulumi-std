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

type Rfc3339tounix struct{}
type Rfc3339tounixargs struct {
	Input string `pulumi:"input"`
}

type Rfc3339tounixresult struct {
	Result int64 `pulumi:"result"`
}

func (r *Rfc3339tounix) Annotate(a infer.Annotator) {
	a.Describe(r, `Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.`)
}

func (*Rfc3339tounix) Call(_ p.Context, args Rfc3339tounixargs) (Rfc3339tounixresult, error) {
	t, err := time.Parse(time.RFC3339, args.Input)
	if err != nil {
		return Rfc3339tounixresult{}, err
	}
	return Rfc3339tounixresult{t.UnixMilli()}, nil
}
