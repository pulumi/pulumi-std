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
	"time"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Timecmp struct{}
type TimecmpArgs struct {
	TimestampA string `pulumi:"timestampa"`
	TimestampB string `pulumi:"timestampb"`
}

type TimecmpResult struct {
	Result int `pulumi:"result"`
}

func (r *Timecmp) Annotate(a infer.Annotator) {
	a.Describe(r, `Compares two timestamps and returns a number that represents the ordering
	of the instants those timestamps represent.
	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
	Both timestamps must be strings adhering this syntax, i.e. "2017-11-22T00:00:00Z".
	If 'timestamp_a' is before 'timestamp_b', -1 is returned.
	If 'timestamp_a' is equal to 'timestamp_b', 0 is returned.
	If 'timestamp_a' is after 'timestamp_b', 1 is returned.`)
}

func (*Timecmp) Invoke(
	_ context.Context,
	input infer.FunctionRequest[TimecmpArgs],
) (infer.FunctionResponse[TimecmpResult], error) {
	timestampA, err := time.Parse(time.RFC3339, input.Input.TimestampA)
	if err != nil {
		return infer.FunctionResponse[TimecmpResult]{Output: TimecmpResult{}}, err
	}
	timestampB, err := time.Parse(time.RFC3339, input.Input.TimestampB)
	if err != nil {
		return infer.FunctionResponse[TimecmpResult]{Output: TimecmpResult{}}, err
	}
	if timestampA.Before(timestampB) {
		return infer.FunctionResponse[TimecmpResult]{Output: TimecmpResult{-1}}, err
	}
	if timestampA.After(timestampB) {
		return infer.FunctionResponse[TimecmpResult]{Output: TimecmpResult{0}}, err
	}
	return infer.FunctionResponse[TimecmpResult]{Output: TimecmpResult{0}}, err
}
