// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Basename
    {
        /// <summary>
        /// Returns the last element of the input path.
        /// </summary>
        public static Task<BasenameResult> InvokeAsync(BasenameArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<BasenameResult>("std:index:basename", args ?? new BasenameArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the last element of the input path.
        /// </summary>
        public static Output<BasenameResult> Invoke(BasenameInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<BasenameResult>("std:index:basename", args ?? new BasenameInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the last element of the input path.
        /// </summary>
        public static Output<BasenameResult> Invoke(BasenameInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<BasenameResult>("std:index:basename", args ?? new BasenameInvokeArgs(), options.WithDefaults());
    }


    public sealed class BasenameArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public BasenameArgs()
        {
        }
        public static new BasenameArgs Empty => new BasenameArgs();
    }

    public sealed class BasenameInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public BasenameInvokeArgs()
        {
        }
        public static new BasenameInvokeArgs Empty => new BasenameInvokeArgs();
    }


    [OutputType]
    public sealed class BasenameResult
    {
        public readonly string Result;

        [OutputConstructor]
        private BasenameResult(string result)
        {
            Result = result;
        }
    }
}
