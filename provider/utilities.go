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
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

func stringHashFunction(hashFn func() hash.Hash, encoder func([]byte) string) func(string) string {
	return func(input string) string {
		h := hashFn()
		h.Write([]byte(input))
		return encoder(h.Sum(nil))
	}
}

func jsonDeepEquals(a, b interface{}) bool {
	aBytes, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bBytes, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return string(aBytes) == string(bBytes)
}

func readFileContents(path string) (string, error) {
	path = filepath.Clean(path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file(%s) failed to read the contents because the file does not exist", path)
		}

		return "", err
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(fileBytes), nil
}
