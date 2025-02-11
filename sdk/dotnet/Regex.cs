// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Regex
    {
        /// <summary>
        /// Returns the first matche of a regular expression in a string (including named or indexed groups).
        /// </summary>
        public static Task<RegexResult> InvokeAsync(RegexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RegexResult>("std:index:regex", args ?? new RegexArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the first matche of a regular expression in a string (including named or indexed groups).
        /// </summary>
        public static Output<RegexResult> Invoke(RegexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RegexResult>("std:index:regex", args ?? new RegexInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the first matche of a regular expression in a string (including named or indexed groups).
        /// </summary>
        public static Output<RegexResult> Invoke(RegexInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<RegexResult>("std:index:regex", args ?? new RegexInvokeArgs(), options.WithDefaults());
    }


    public sealed class RegexArgs : global::Pulumi.InvokeArgs
    {
        [Input("pattern", required: true)]
        public string Pattern { get; set; } = null!;

        [Input("string", required: true)]
        public string String { get; set; } = null!;

        public RegexArgs()
        {
        }
        public static new RegexArgs Empty => new RegexArgs();
    }

    public sealed class RegexInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("string", required: true)]
        public Input<string> String { get; set; } = null!;

        public RegexInvokeArgs()
        {
        }
        public static new RegexInvokeArgs Empty => new RegexInvokeArgs();
    }


    [OutputType]
    public sealed class RegexResult
    {
        public readonly object Result;

        [OutputConstructor]
        private RegexResult(object result)
        {
            Result = result;
        }
    }
}
