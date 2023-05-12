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

type Timestamp struct{}
type TimestampArgs struct{}

type TimestampResult struct {
	Result string `pulumi:"result"`
}

func (r *Timestamp) Annotate(a infer.Annotator) {
	a.Describe(r, `Returns a UTC timestamp string of the current time in RFC 3339 format`)
}

func (*Timestamp) Call(_ p.Context, _ TimestampArgs) (TimestampResult, error) {
	return TimestampResult{time.Now().Format(time.RFC3339)}, nil
}
