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
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"go": map[string]any{
					"importBasePath": "github.com/pulumi/pulumi-std/sdk/go/std",
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
			infer.Function[*Abspath, AbspathArgs, AbspathResult](),
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
			infer.Function[*Csvdecode, CsvdecodeArgs, CsvdecodeResult](),
			infer.Function[*Dirname, DirnameArgs, DirnameResult](),
			infer.Function[*Distinct, DistinctArgs, DistinctResult](),
			infer.Function[*Element, ElementArgs, ElementResult](),
			infer.Function[*File, FileArgs, FileResult](),
			infer.Function[*Fileexists, FileexistsArgs, FileexistsResult](),
			infer.Function[*Filemd5, Filemd5Args, Filemd5Result](),
			infer.Function[*Filesha1, Filesha1Args, Filesha1Result](),
			infer.Function[*Filesha256, Filesha256Args, Filesha256Result](),
			infer.Function[*Filesha512, Filesha512Args, Filesha512Result](),
			infer.Function[*Filebase64, Filebase64Args, Filebase64Result](),
			infer.Function[*Filebase64sha256, Filebase64sha256Args, Filebase64sha256Result](),
			infer.Function[*Filebase64sha512, Filebase64sha512Args, Filebase64sha512Result](),
			infer.Function[*Floor, FloorArgs, FloorResult](),
			infer.Function[*Flatten, FlattenArgs, FlattenResult](),
			infer.Function[*Format, FormatArgs, FormatResult](),
			infer.Function[*Index, IndexArgs, IndexResult](),
			infer.Function[*Indent, IndentArgs, IndentResult](),
			infer.Function[*Join, JoinArgs, JoinResult](),
			infer.Function[*Jsonencode, JsonencodeArgs, JsonencodeResult](),
			infer.Function[*Jsondecode, JsondecodeArgs, JsondecodeResult](),
			infer.Function[*Keys, KeysArgs, KeysResult](),
			infer.Function[*Length, LengthArgs, LengthResult](),
			infer.Function[*Log, LogArgs, LogResult](),
			infer.Function[*Lookup, LookupArgs, LookupResult](),
			infer.Function[*Lower, LowerArgs, LowerResult](),
			infer.Function[*Map, MapArgs, MapResult](),
			infer.Function[*Max, MaxArgs, MaxResult](),
			infer.Function[*Md5, Md5Args, Md5Result](),
			infer.Function[*Matchkeys, MatchkeysArgs, MatchkeysResult](),
			infer.Function[*Merge, MergeArgs, MergeResult](),
			infer.Function[*Min, MinArgs, MinResult](),
			infer.Function[*Pathexpand, PathexpandArgs, PathexpandResult](),
			infer.Function[*Pow, PowArgs, PowResult](),
			infer.Function[*Rsadecrypt, RsadecryptArgs, RsadecryptResult](),
			infer.Function[*Sha1, Sha1Args, Sha1Result](),
			infer.Function[*Sha256, Sha256Args, Sha256Result](),
			infer.Function[*Sha512, Sha512Args, Sha512Result](),
			infer.Function[*Signum, SignumArgs, SignumResult](),
			infer.Function[*Slice, SliceArgs, SliceResult](),
			infer.Function[*Parseint, ParseintArgs, ParseintResult](),
			infer.Function[*Reverse, ReverseArgs, ReverseResult](),
			infer.Function[*Sort, SortArgs, SortResult](),
			infer.Function[*Startswith, StartswithArgs, StartswithResult](),
			infer.Function[*Endswith, EndswithArgs, EndswithResult](),
			infer.Function[*Values, ValuesArgs, ValuesResult](),
			infer.Function[*Range, RangeArgs, RangeResults](),
			infer.Function[*Replace, ReplaceArgs, ReplaceResult](),
			infer.Function[*Split, SplitArgs, SplitResult](),
			infer.Function[*Strrev, StrrevArgs, StrrevResult](),
			infer.Function[*Substr, SubstrArgs, SubstrResult](),
			infer.Function[*Title, TitleArgs, TitleResult](),
			infer.Function[*Timeadd, TimeaddArgs, TimeaddResult](),
			infer.Function[*Timecmp, TimecmpArgs, TimecmpResult](),
			infer.Function[*Timestamp, TimestampArgs, TimestampResult](),
			infer.Function[*Rfc3339tounix, Rfc3339tounixargs, Rfc3339tounixresult](),
			infer.Function[*Unixtorfc3999, Unixtorfc3999args, Unixtorfc3999result](),
			infer.Function[*Tobool, ToboolArgs, ToboolResult](),
			infer.Function[*Tolist, TolistArgs, TolistResult](),
			infer.Function[*Toset, TosetArgs, TosetResult](),
			infer.Function[*Tonumber, TonumberArgs, TonumberResult](),
			infer.Function[*Tostring, TostringArgs, TostringResult](),
			infer.Function[*Transpose, TransposeArgs, TransposeResult](),
			infer.Function[*Trim, TrimArgs, TrimResult](),
			infer.Function[*Trimprefix, TrimprefixArgs, TrimprefixResult](),
			infer.Function[*Trimsuffix, TrimsuffixArgs, TrimsuffixResult](),
			infer.Function[*Trimspace, TrimspaceArgs, TrimspaceResult](),
			infer.Function[*Upper, UpperArgs, UpperResult](),
			infer.Function[*Urlencode, UrlencodeArgs, UrlencodeResult](),
			infer.Function[*Uuid, UuidArgs, UuidResult](),
			infer.Function[*Alltrue, AlltrueArgs, AlltrueResult](),
			infer.Function[*Anytrue, AnytrueArgs, AnytrueResult](),
			infer.Function[*Sum, SumArgs, SumResult](),
			infer.Function[*Zipmap, ZipmapArgs, ZipmapResult](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"std": "index",
		},
	})
}
