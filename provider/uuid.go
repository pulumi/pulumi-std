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

	"github.com/google/uuid"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// nolint
type Uuid struct{}

// nolint
type UuidArgs struct{}

// nolint
type UuidResult struct {
	Result string `pulumi:"result"`
}

func (r *Uuid) Annotate(a infer.Annotator) {
	a.Describe(r, "Returns a unique identifier string, generated and formatted as required by RFC 4122.")
}

func (*Uuid) Invoke(
	_ context.Context,
	_ infer.FunctionRequest[UuidArgs],
) (infer.FunctionResponse[UuidResult], error) {
	return infer.FunctionResponse[UuidResult]{Output: UuidResult{uuid.New().String()}}, nil
}
