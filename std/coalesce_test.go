// Copyright 2025, Pulumi Corporation.
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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoalesceAssignableType(t *testing.T) {
	t.Parallel()

	// Arrange.
	boolT := reflect.TypeOf(true)
	int32T := reflect.TypeOf(int32(0))
	int64T := reflect.TypeOf(int64(0))
	float64T := reflect.TypeOf(float64(0.0))
	stringT := reflect.TypeOf("")

	cases := []struct {
		name          string
		input         []interface{}
		expected      reflect.Type
		expectedError string
	}{
		{
			name:     "All small integers",
			input:    []interface{}{1, 2, 3},
			expected: int32T,
		},
		{
			name:     "All big integers",
			input:    []interface{}{10_000_000_000, 20_000_000_000},
			expected: int64T,
		},
		{
			name:     "Big and small integers 1",
			input:    []interface{}{1, 2, 50_000_000_000, 3, 10_000_000_000},
			expected: int64T,
		},
		{
			name:     "Big and small integers 2",
			input:    []interface{}{50_000_000_000, 1, 2, 3, 10_000_000_000},
			expected: int64T,
		},
		{
			name:     "Big and small integers 2",
			input:    []interface{}{50_000_000_000, 1, 2, 3, 10_000_000_000, 4},
			expected: int64T,
		},
		{
			name:          "Big integers and floats 1",
			input:         []interface{}{1.0, 2, 50_000_000_000, 3, 10_000_000_000},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Big integers and floats 2",
			input:         []interface{}{50_000_000_000, 3.141, 10_000_000_000},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:     "All floats",
			input:    []interface{}{1.0, 2.0, 3.0},
			expected: float64T,
		},
		{
			name:     "Small integers and floats 1",
			input:    []interface{}{1, 2, 3.0},
			expected: float64T,
		},
		{
			name:     "Small integers and floats 2",
			input:    []interface{}{1.0, 2, 3},
			expected: float64T,
		},
		{
			name:     "Small integers and floats 2",
			input:    []interface{}{1.0, 2, 3.0},
			expected: float64T,
		},
		{
			name:     "Integers and strings 1",
			input:    []interface{}{1, 2, "3"},
			expected: stringT,
		},
		{
			name:     "Integers and strings 2",
			input:    []interface{}{"1", 2, 3},
			expected: stringT,
		},
		{
			name:     "Integers and strings 3",
			input:    []interface{}{"1", 2, "3"},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 1",
			input:    []interface{}{1, 2, "3", 4.0},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 2",
			input:    []interface{}{"1", 2, "3", 4.0},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 3",
			input:    []interface{}{"1", 2.0, "3", 4},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 4",
			input:    []interface{}{10_000_000_000, 2.0, "3", 4, 5.0},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 5",
			input:    []interface{}{"3", 2.0, "3", 4, 5.0, 6},
			expected: stringT,
		},
		{
			name:     "Integers, strings, and floats 6",
			input:    []interface{}{"3", 20_000_000_000, "3", 4, 5.0, 6},
			expected: stringT,
		},
		{
			name:     "Integers, strings, floats, and nil",
			input:    []interface{}{1, 2, "3", 4.0, nil},
			expected: stringT,
		},
		{
			name:     "Integers, strings, floats, nil, and more",
			input:    []interface{}{1, 2, "3", 4.0, nil, 5},
			expected: stringT,
		},
		{
			name:     "All strings",
			input:    []interface{}{"1", "2", "3"},
			expected: stringT,
		},
		{
			name:     "All booleans",
			input:    []interface{}{true, false, true},
			expected: boolT,
		},
		{
			name:     "Booleans and strings",
			input:    []interface{}{false, true, "foo"},
			expected: stringT,
		},
		{
			name:     "Booleans, strings, integers and floats",
			input:    []interface{}{false, true, "3", 4, 5.0, 50_000_000_000},
			expected: stringT,
		},
		{
			name:          "Small integers and booleans 1",
			input:         []interface{}{1, true},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Small integers and booleans 2",
			input:         []interface{}{false, 2},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Big integers and booleans 1",
			input:         []interface{}{50_000_000_000, false},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Big integers and booleans 2",
			input:         []interface{}{true, 50_000_000_000},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Floats and booleans 1",
			input:         []interface{}{1.0, false},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:          "Floats and booleans 2",
			input:         []interface{}{true, 2.0},
			expectedError: "all arguments must be convertible to the same type",
		},
		{
			name:     "All nil",
			input:    []interface{}{nil, nil, nil},
			expected: nil,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// Act.
			actual, err := assignableType(c.input)

			// Assert.
			if c.expectedError != "" {
				assert.ErrorContains(t, err, c.expectedError)
				assert.Nil(t, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.expected, actual, "mostGenericType(%v) = %v; want %v", c.input, actual, c.expected)
			}
		})
	}
}
