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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func exitOnError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "\n", err)
		cleanup()
		os.Exit(1)
	}
}

func stateDir() string {
	workingDir, err := os.Getwd()
	exitOnError("Could not get working directory", err)
	return fmt.Sprintf("%s/state", workingDir)
}

func grpcJsonDebug() string {
	workingDir, err := os.Getwd()
	exitOnError("Could not get working directory", err)
	return fmt.Sprintf("%s/grpc.json", workingDir)
}

func prepare() {
	// local backend requires a passphrase
	err := os.Setenv("PULUMI_CONFIG_PASSPHRASE", "test")
	exitOnError("Could not set PULUMI_CONFIG_PASSPHRASE", err)
	err = os.Setenv("PULUMI_DEBUG_GRPC", grpcJsonDebug())
	stateDir := stateDir()
	// make state dir
	err = os.Mkdir(stateDir, 0755)
	exitOnError("Could not create state directory", err)
	_, err = exec.Command("pulumi", "login", "--cloud-url", fmt.Sprintf("file://%s", stateDir)).Output()
	_, err = exec.Command("pulumi", "stack", "init", "dev").Output()
	exitOnError("Could not init pulumi stack", err)
}

func cleanup() {
	_, err := exec.Command("pulumi", "logout").Output()
	exitOnError("Could not logout from pulumi", err)
	stateDir := stateDir()
	err = os.RemoveAll(stateDir)
	exitOnError("Could not remove state directory", err)
	err = os.RemoveAll("Pulumi.dev.yaml")
	exitOnError("Could not remove Pulumi.dev.yaml", err)
}

func outputs() map[string]interface{} {
	up, err := exec.Command("pulumi", "up", "--stack", "dev", "--yes").Output()
	fmt.Println(string(up))
	exitOnError("Could not run pulumi up", err)
	stdout, err := exec.Command("pulumi", "stack", "output", "--json").Output()
	exitOnError("Could not get pulumi stack outputs", err)
	var outputs map[string]interface{}
	err = json.Unmarshal(stdout, &outputs)
	exitOnError("Could not unmarshal pulumi stack outputs", err)
	return outputs
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

func expectedOutputs() map[string]interface{} {
	cwd, _ := os.Getwd()

	return map[string]interface{}{
		"abs-positive":               42.24,
		"abs-negative":               42.24,
		"abspath":                    fmt.Sprintf("%s/myfile.txt", cwd),
		"base64encode":               "QUJDREU=",
		"base64decode":               "ABCDE",
		"basename":                   "pulumi",
		"base64gzip":                 "H4sIAAAAAAAA/0pMSgYAAAD//wEAAP//wkEkNQMAAAA=",
		"base64sha256":               "ungWv48Bz+pBQUDeXa4iI7ADYaOWF3qctBD/YfIAFa0=",
		"base64sha512":               "3a81oZNherrMQXNJriBBMRLm+k6JqX6iCp7u5ktV05ohkpkqJ0/BqDa6PCOj/uu9RU1EI2Q86A4qmslPpUyknw==",
		"chunklist":                  []interface{}{[]int{1, 2}, []int{3, 4}, []int{5}},
		"chunklistByOne":             []interface{}{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}},
		"cidrhost":                   "10.0.0.2",
		"cidrhostV2":                 "10.255.255.254",
		"cidrnetmask":                "255.0.0.0",
		"cidrsubnet":                 "10.2.0.0/16",
		"coalesce":                   "hello",
		"coalescelist":               []interface{}{1},
		"compact":                    []string{"one"},
		"concat":                     []int{1, 2, 3, 4, 5},
		"containsWithIntegers":       true,
		"containsWithStrings":        true,
		"containsWithLists":          true,
		"csvdecode":                  []map[string]string{{"a": "1", "b": "2", "c": "3"}, {"a": "4", "b": "5", "c": "6"}},
		"shouldNotContain":           false,
		"dirname":                    "/path/to/directory",
		"distinctWithIntegers":       []int{1, 2, 3, 4, 5},
		"distinctWithStrings":        []string{"one", "two", "three"},
		"element":                    20,
		"elementOverflow":            10,
		"elementNegativeIndex":       30,
		"file":                       "hello world",
		"fileExistsTrue":             true,
		"filemd5":                    "5eb63bbbe01eeed093cb22bb8f5acdc3",
		"filesha1":                   "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
		"filesha256":                 "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		"filesha512":                 "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f",
		"filebase64":                 "aGVsbG8gd29ybGQ=",
		"filebase64sha256":           "Yjk0ZDI3Yjk5MzRkM2UwOGE1MmU1MmQ3ZGE3ZGFiZmFjNDg0ZWZlMzdhNTM4MGVlOTA4OGY3YWNlMmVmY2RlOQ==",
		"filebase64sha512":           "MzA5ZWNjNDg5YzEyZDZlYjRjYzQwZjUwYzkwMmYyYjRkMGVkNzdlZTUxMWE3YzdhOWJjZDNjYTg2ZDRjZDg2Zjk4OWRkMzViYzVmZjQ5OTY3MGRhMzQyNTViNDViMGNmZDgzMGU4MWY2MDVkY2Y3ZGM1NTQyZTkzYWU5Y2Q3NmY=",
		"floor":                      1.0,
		"flatten":                    []int{1, 2, 3, 4, 5},
		"flattenWithStrings":         []string{"one", "two", "three", "four", "five"},
		"format":                     "hello-world",
		"index":                      1,
		"join":                       "one-two-three",
		"jsonencode":                 "{\"hello\":\"world\"}",
		"jsondecode":                 map[string]interface{}{"str": "hello", "num": 6.13, "strs": []string{"a", "b"}, "bool": true},
		"keys":                       []string{"hello", "one"},
		"log":                        3.0,
		"md5":                        "900150983cd24fb0d6963f7d28e17f72",
		"slice":                      []int{1, 2, 3},
		"parseint":                   100,
		"parseintWithBase":           255,
		"reverse":                    []interface{}{"foo", 1, true},
		"range":                      []float64{0, 1, 2},
		"rangeStartLimit":            []float64{1, 2, 3},
		"rangeStartLimitStep":        []float64{1, 3, 5, 7},
		"rangeStartLimitStepDecimal": []float64{1, 1.5, 2, 2.5, 3, 3.5},
		"rangeStartLimitDecline":     []float64{4, 3, 2},
		"rangeStartLimitStepDecline": []float64{10, 8, 6},
		"sort":                       []string{"Avocado", "Orange", "apple", "banana", "orange", "pear", "watermelon"},
		"startswith":                 true,
		"endswith":                   true,
		"values":                     []int{1, 2},
		"split":                      []string{"one", "two", "three"},
		"strrev":                     "olleh",
		"substr":                     "hello",
		"transpose":                  map[string][]string{"1": {"a"}, "2": {"a", "b"}, "3": {"b"}},
		"title":                      "Hello World",
		"toboolBool":                 true,
		"toboolStr":                  false,
		"tolist":                     []string{"a", "true", "3.14"},
		"toset":                      []string{"a", "true", "3.14"},
		"tonumberInt":                3,
		"tonumberFloat":              3.14159,
		"tonumberStr":                10,
		"tostringStr":                "hello world",
		"tostringInt":                "8",
		"tostringFloat":              "1.234",
		"tostringBool":               "true",
		"trim":                       "Hello  World!",
		"trimprefix":                 " bar",
		"trimsuffix":                 "foo ",
		"trimspace":                  "foobar",
		"upper":                      "ONE",
		"urlencode":                  "Hello+World%21",
		"alltrue":                    true,
		"anytrue":                    true,
		"sum":                        15,
		"zipmap":                     map[string]interface{}{"one": 1, "two": 2, "three": 3},
	}
}

func main() {
	cleanup()
	prepare()

	exitCode := 0
	exit := func() {
		cleanup()
		os.Exit(exitCode)
	}

	defer func() {
		// If we panic, we want to make sure we clean up.
		if r := recover(); r != nil {
			exitCode = 1
			exit()
		}
	}()

	defer exit()

	expected := expectedOutputs()
	outputs := outputs()

	for key, value := range outputs {

		if key == "pathexpand" {
			// pathexpand is a special case, because it returns a different value on each machine
			if strings.Contains(value.(string), "~") {
				fmt.Printf("❌ Output '%s' => '%s' did not expand the home directory", key, value)
				exitCode++
			}
		}

		expectedValue, ok := expected[key]
		if !ok {
			continue
		}

		if reflect.DeepEqual(expectedValue, value) || jsonDeepEquals(expectedValue, value) {
			fmt.Printf("✅ Output '%s' has value '%v' as expected\n", key, value)
		} else {
			fmt.Printf("❌ Output '%s' was '%v' but should be '%v'\n", key, value, expectedValue)
			exitCode++
		}
	}
}
