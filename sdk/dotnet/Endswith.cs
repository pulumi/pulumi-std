// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Endswith
    {
        /// <summary>
        /// Determines if the input string ends with the suffix.
        /// </summary>
        public static Task<EndswithResult> InvokeAsync(EndswithArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<EndswithResult>("std:index:endswith", args ?? new EndswithArgs(), options.WithDefaults());

        /// <summary>
        /// Determines if the input string ends with the suffix.
        /// </summary>
        public static Output<EndswithResult> Invoke(EndswithInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<EndswithResult>("std:index:endswith", args ?? new EndswithInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Determines if the input string ends with the suffix.
        /// </summary>
        public static Output<EndswithResult> Invoke(EndswithInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<EndswithResult>("std:index:endswith", args ?? new EndswithInvokeArgs(), options.WithDefaults());
    }


    public sealed class EndswithArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        [Input("suffix", required: true)]
        public string Suffix { get; set; } = null!;

        public EndswithArgs()
        {
        }
        public static new EndswithArgs Empty => new EndswithArgs();
    }

    public sealed class EndswithInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public EndswithInvokeArgs()
        {
        }
        public static new EndswithInvokeArgs Empty => new EndswithInvokeArgs();
    }


    [OutputType]
    public sealed class EndswithResult
    {
        public readonly bool Result;

        [OutputConstructor]
        private EndswithResult(bool result)
        {
            Result = result;
        }
    }
}
