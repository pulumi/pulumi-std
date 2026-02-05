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
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			Description: "Standard library functions",
			DisplayName: "StandardLibrary",
			Publisher:   "Pulumi",
			Homepage:    "https://github.com/pulumi/pulumi-std",
			Repository:  "https://github.com/pulumi/pulumi-std",
			LanguageMap: map[string]any{
				"go": map[string]any{
					"respectSchemaVersion": true,
					"importBasePath":       "github.com/pulumi/pulumi-std/sdk/v2/go/std",
				},
				"java": map[string]any{
					"buildFiles":                      "gradle",
					"gradleNexusPublishPluginVersion": "2.0.0",
				},
				"csharp": map[string]any{
					"respectSchemaVersion": true,
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
				"nodejs": map[string]any{
					"respectSchemaVersion": true,
				},
				"python": map[string]any{
					"pyproject": map[string]any{
						"enabled": true,
					},
					"respectSchemaVersion": true,
				},
			},
		},
		Functions: []infer.InferredFunction{
			infer.Function(&Abs{}),
			infer.Function(&Abspath{}),
			infer.Function(&Basename{}),
			infer.Function(&Base64decode{}),
			infer.Function(&Base64encode{}),
			infer.Function(&Base64gzip{}),
			infer.Function(&Base64sha256{}),
			infer.Function(&Base64sha512{}),
			infer.Function(&Bcrypt{}),
			infer.Function(&Ceil{}),
			infer.Function(&Chomp{}),
			infer.Function(&Chunklist{}),
			infer.Function(&Cidrhost{}),
			infer.Function(&Cidrnetmask{}),
			infer.Function(&Cidrsubnet{}),
			infer.Function(&Cidrsubnets{}),
			infer.Function(&Coalesce{}),
			infer.Function(&Coalescelist{}),
			infer.Function(&Compact{}),
			infer.Function(&Concat{}),
			infer.Function(&Contains{}),
			infer.Function(&Csvdecode{}),
			infer.Function(&Dirname{}),
			infer.Function(&Distinct{}),
			infer.Function(&Element{}),
			infer.Function(&File{}),
			infer.Function(&Fileexists{}),
			infer.Function(&Filemd5{}),
			infer.Function(&Filesha1{}),
			infer.Function(&Filesha256{}),
			infer.Function(&Filesha512{}),
			infer.Function(&Filebase64{}),
			infer.Function(&Filebase64sha256{}),
			infer.Function(&Filebase64sha512{}),
			infer.Function(&Floor{}),
			infer.Function(&Flatten{}),
			infer.Function(&Format{}),
			infer.Function(&Formatlist{}),
			infer.Function(&Index{}),
			infer.Function(&Indent{}),
			infer.Function(&Join{}),
			infer.Function(&Jsonencode{}),
			infer.Function(&Jsondecode{}),
			infer.Function(&Keys{}),
			infer.Function(&Length{}),
			infer.Function(&Log{}),
			infer.Function(&Lookup{}),
			infer.Function(&Lower{}),
			infer.Function(&Map{}),
			infer.Function(&Max{}),
			infer.Function(&Md5{}),
			infer.Function(&Matchkeys{}),
			infer.Function(&Merge{}),
			infer.Function(&Min{}),
			infer.Function(&Pathexpand{}),
			infer.Function(&Pow{}),
			infer.Function(&Rsadecrypt{}),
			infer.Function(&Sha1{}),
			infer.Function(&Sha256{}),
			infer.Function(&Sha512{}),
			infer.Function(&Signum{}),
			infer.Function(&Slice{}),
			infer.Function(&Parseint{}),
			infer.Function(&Reverse{}),
			infer.Function(&Sort{}),
			infer.Function(&Startswith{}),
			infer.Function(&Endswith{}),
			infer.Function(&Values{}),
			infer.Function(&Setintersection{}),
			infer.Function(&Range{}),
			infer.Function(&Replace{}),
			infer.Function(&Regex{}),
			infer.Function(&Regexall{}),
			infer.Function(&Split{}),
			infer.Function(&Strrev{}),
			infer.Function(&Substr{}),
			infer.Function(&Title{}),
			infer.Function(&Timeadd{}),
			infer.Function(&Timecmp{}),
			infer.Function(&Timestamp{}),
			infer.Function(&Tobool{}),
			infer.Function(&Tolist{}),
			infer.Function(&Toset{}),
			infer.Function(&Tonumber{}),
			infer.Function(&Tostring{}),
			infer.Function(&Transpose{}),
			infer.Function(&Trim{}),
			infer.Function(&Trimprefix{}),
			infer.Function(&Trimsuffix{}),
			infer.Function(&Trimspace{}),
			infer.Function(&Upper{}),
			infer.Function(&Urlencode{}),
			infer.Function(&Uuid{}),
			infer.Function(&Alltrue{}),
			infer.Function(&Anytrue{}),
			infer.Function(&Sum{}),
			infer.Function(&Zipmap{}),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"v2": "index",
		},
	})
}
