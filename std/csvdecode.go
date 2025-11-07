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
	"errors"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Csvdecode struct{}
type CsvdecodeArgs struct {
	Input string `pulumi:"input"`
}

type CsvdecodeResult struct {
	Result []map[string]string `pulumi:"result"`
}

func (r *Csvdecode) Annotate(a infer.Annotator) {
	a.Describe(r, `Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
	The first line of the CSV data is interpreted as a "header" row: the values given
	are used as the keys in the resulting maps.
	Each subsequent line becomes a single map in the resulting list,
	matching the keys from the header row with the given values by index.
	All lines in the file must contain the same number of fields,
	or this function will produce an error.
	Follows the format defined in RFC 4180.`)
}

func (*Csvdecode) Call(_ context.Context, args CsvdecodeArgs) (CsvdecodeResult, error) {
	res := make([]map[string]string, 0)
	if args.Input == "" {
		return CsvdecodeResult{res}, nil
	}
	rows := strings.Split(args.Input, "\n")
	if len(rows) == 1 {
		return CsvdecodeResult{res}, nil
	}

	header := strings.Split(rows[0], ",")
	numColumns := len(header)
	for i := 1; i < len(rows); i++ {
		row := strings.Split(rows[i], ",")
		if len(row) != numColumns {
			return CsvdecodeResult{res}, errors.New("Invalid input: each line must contain the same number of fields")
		}
		m := make(map[string]string)
		for c := 0; c < numColumns; c++ {
			k, v := strings.TrimSpace(header[c]), strings.TrimSpace(row[c])
			m[k] = v
		}
		res = append(res, m)
	}

	return CsvdecodeResult{res}, nil
}
