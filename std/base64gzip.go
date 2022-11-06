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
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Base64gzip struct{}
type Base64GzipArgs struct {
	Input string `pulumi:"input"`
}

type Base64GzipResult struct {
	Result string `pulumi:"result"`
}

func (r *Base64gzip) Annotate(a infer.Annotator) {
	a.Describe(r, `Compresses the given string with gzip and then encodes the result to base64.`)
}

func base64gzip(input string) (string, error) {
	var buffer bytes.Buffer
	gz := gzip.NewWriter(&buffer)
	if _, err := gz.Write([]byte(input)); err != nil {
		return "", fmt.Errorf("failed to write gzip raw data: %w", err)
	}
	if err := gz.Flush(); err != nil {
		return "", fmt.Errorf("failed to flush gzip writer: %w", err)
	}
	if err := gz.Close(); err != nil {
		return "", fmt.Errorf("failed to close gzip writer: %w", err)
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func (*Base64gzip) Call(ctx p.Context, args Base64GzipArgs) (Base64GzipResult, error) {
	result, err := base64gzip(args.Input)
	if err != nil {
		return Base64GzipResult{}, err
	}
	return Base64GzipResult{result}, nil
}
