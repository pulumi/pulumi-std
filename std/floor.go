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

type Floor struct{}
type FloorArgs struct {
	Input float64 `pulumi:"input"`
}

type FloorResult struct {
	Result float64 `pulumi:"result"`
}

func (r *Floor) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns the greatest integer value less than or equal to the argument.")
}

func (*Floor) Call(_ p.Context, input FloorArgs) (FloorResult, error) {
	return FloorResult{math.Floor(input.Input)}, nil
}
