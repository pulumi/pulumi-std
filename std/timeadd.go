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

type Timeadd struct{}
type TimeaddArgs struct {
	Timestamp string `pulumi:"timestamp"`
	Duration  string `pulumi:"duration"`
}

type TimeaddResult struct {
	Result string `pulumi:"result"`
}

func (r *Timeadd) Annotate(a infer.Annotator) {
	a.Describe(r, `Adds a duration to a timestamp, returning a new timestamp.
	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
	'timestamp' must be a string adhering this syntax, i.e. "2017-11-22T00:00:00Z".
	'duration' is a string representation of a time difference, comprised of sequences of
	numbers and unit pairs, i.e. "3.5h" or "2h15m".
	Accepted units are "ns", "us" or "µs", "ms", "s", "m", and "h". The first number may be negative
	to provide a negative duration, i.e. "-2h15m".`)
}

func (*Timeadd) Call(_ context.Context, args TimeaddArgs) (TimeaddResult, error) {
	duration, err := time.ParseDuration(args.Duration)
	if err != nil {
		return TimeaddResult{}, err
	}
	timestamp, err := time.Parse(time.RFC3339, args.Timestamp)
	if err != nil {
		return TimeaddResult{}, err
	}
	return TimeaddResult{timestamp.Add(duration).Format(time.RFC3339)}, nil
}
