// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Csvdecode
    {
        /// <summary>
        /// Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
        /// 	The first line of the CSV data is interpreted as a "header" row: the values given 
        /// 	are used as the keys in the resulting maps. 
        /// 	Each subsequent line becomes a single map in the resulting list, 
        /// 	matching the keys from the header row with the given values by index. 
        /// 	All lines in the file must contain the same number of fields, 
        /// 	or this function will produce an error.
        /// 	Follows the format defined in RFC 4180.
        /// </summary>
        public static Task<CsvdecodeResult> InvokeAsync(CsvdecodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CsvdecodeResult>("std:index:csvdecode", args ?? new CsvdecodeArgs(), options.WithDefaults());

        /// <summary>
        /// Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
        /// 	The first line of the CSV data is interpreted as a "header" row: the values given 
        /// 	are used as the keys in the resulting maps. 
        /// 	Each subsequent line becomes a single map in the resulting list, 
        /// 	matching the keys from the header row with the given values by index. 
        /// 	All lines in the file must contain the same number of fields, 
        /// 	or this function will produce an error.
        /// 	Follows the format defined in RFC 4180.
        /// </summary>
        public static Output<CsvdecodeResult> Invoke(CsvdecodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CsvdecodeResult>("std:index:csvdecode", args ?? new CsvdecodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class CsvdecodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public CsvdecodeArgs()
        {
        }
        public static new CsvdecodeArgs Empty => new CsvdecodeArgs();
    }

    public sealed class CsvdecodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public CsvdecodeInvokeArgs()
        {
        }
        public static new CsvdecodeInvokeArgs Empty => new CsvdecodeInvokeArgs();
    }


    [OutputType]
    public sealed class CsvdecodeResult
    {
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Result;

        [OutputConstructor]
        private CsvdecodeResult(ImmutableArray<ImmutableDictionary<string, string>> result)
        {
            Result = result;
        }
    }
}
