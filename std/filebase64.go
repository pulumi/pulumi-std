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
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Filebase64 struct{}
type Filebase64Args struct {
	Input string `pulumi:"input"`
}

type Filebase64Result struct {
	Result string `pulumi:"result"`
}

func (r *Filebase64) Annotate(a infer.Annotator) {
	a.Describe(r, "Reads the contents of a file and returns them as a base64-encoded string.")
}

func readFileContentsToBase64(path string) (string, error) {
	path = filepath.Clean(path)
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file(%s) failed to read the contents because the file does not exist", path)
		}

		return "", err
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return base64.StdEncoding.EncodeToString(fileBytes), nil
}

func (*Filebase64) Call(ctx p.Context, args Filebase64Args) (Filebase64Result, error) {
	encoded, err := readFileContentsToBase64(args.Input)
	return Filebase64Result{encoded}, err
}
