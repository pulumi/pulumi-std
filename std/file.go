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
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"io"
	"os"
	"path/filepath"
)

type File struct{}
type FileArgs struct {
	Input string `pulumi:"input"`
}

type FileResult struct {
	Result string `pulumi:"result"`
}

func (r *File) Annotate(a infer.Annotator) {
	a.Describe(r, "Reads the contents of a file into the string.")
}

func readFileContents(path string) (string, error) {
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

	return string(fileBytes), nil
}

func (*File) Call(ctx p.Context, args FileArgs) (FileResult, error) {
	contents, err := readFileContents(args.Input)
	if err != nil {
		return FileResult{}, err
	}
	return FileResult{contents}, nil
}
