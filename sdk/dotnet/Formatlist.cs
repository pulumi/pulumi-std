// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Formatlist
    {
        /// <summary>
        /// Formats a list of strings according to the given format. Argument values which are lists are "zipped" together to produce a list of results.
        /// </summary>
        public static Task<FormatlistResult> InvokeAsync(FormatlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<FormatlistResult>("std:index:formatlist", args ?? new FormatlistArgs(), options.WithDefaults());

        /// <summary>
        /// Formats a list of strings according to the given format. Argument values which are lists are "zipped" together to produce a list of results.
        /// </summary>
        public static Output<FormatlistResult> Invoke(FormatlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<FormatlistResult>("std:index:formatlist", args ?? new FormatlistInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Formats a list of strings according to the given format. Argument values which are lists are "zipped" together to produce a list of results.
        /// </summary>
        public static Output<FormatlistResult> Invoke(FormatlistInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<FormatlistResult>("std:index:formatlist", args ?? new FormatlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class FormatlistArgs : global::Pulumi.InvokeArgs
    {
        [Input("args", required: true)]
        private List<object>? _args;
        public List<object> Args
        {
            get => _args ?? (_args = new List<object>());
            set => _args = value;
        }

        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public FormatlistArgs()
        {
        }
        public static new FormatlistArgs Empty => new FormatlistArgs();
    }

    public sealed class FormatlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("args", required: true)]
        private InputList<object>? _args;
        public InputList<object> Args
        {
            get => _args ?? (_args = new InputList<object>());
            set => _args = value;
        }

        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public FormatlistInvokeArgs()
        {
        }
        public static new FormatlistInvokeArgs Empty => new FormatlistInvokeArgs();
    }


    [OutputType]
    public sealed class FormatlistResult
    {
        public readonly ImmutableArray<string> Result;

        [OutputConstructor]
        private FormatlistResult(ImmutableArray<string> result)
        {
            Result = result;
        }
    }
}
