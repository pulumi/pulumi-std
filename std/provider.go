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
				"nodejs": map[string]any{
					"respectSchemaVersion": true,
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
			},
		},
		Functions: []infer.InferredFunction{
			infer.Function[*Abs, AbsArgs, AbsResult](),
			infer.Function[*Basename, BasenameArgs, BasenameResult](),
			infer.Function[*Base64decode, Base64DecodeArgs, Base64DecodeResult](),
			infer.Function[*Base64encode, Base64EncodeArgs, Base64EncodeResult](),
			infer.Function[*Base64gzip, Base64GzipArgs, Base64GzipResult](),
			infer.Function[*Base64sha256, Base64Sha256Args, Base64Sha256Result](),
			infer.Function[*Base64sha512, Base64Sha512Args, Base64Sha512Result](),
			infer.Function[*Bcrypt, BcryptArgs, BcryptResult](),
			infer.Function[*Ceil, CeilArgs, CeilResult](),
			infer.Function[*Chomp, ChompArgs, ChompResult](),
			infer.Function[*Chunklist, ChunklistArgs, ChunklistResult](),
			infer.Function[*Cidrhost, CidrhostArgs, CidrhostResult](),
			infer.Function[*Cidrnetmask, CidrnetmaskArgs, CidrnetmaskResult](),
			infer.Function[*Cidrsubnet, CidrsubnetArgs, CidrsubnetResult](),
			infer.Function[*Coalesce, CoalesceArgs, CoalesceResult](),
			infer.Function[*Coalescelist, CoalescelistArgs, CoalescelistResult](),
			infer.Function[*Compact, CompactArgs, CompactResult](),
			infer.Function[*Concat, ConcatArgs, ConcatResult](),
			infer.Function[*Contains, ContainsArgs, ContainsResult](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"std": "index",
		},
	})
}
