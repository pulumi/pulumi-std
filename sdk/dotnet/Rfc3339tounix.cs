// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Rfc3339tounix
    {
        /// <summary>
        /// Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
        /// </summary>
        public static Task<Rfc3339tounixResult> InvokeAsync(Rfc3339tounixArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Rfc3339tounixResult>("std:index:rfc3339tounix", args ?? new Rfc3339tounixArgs(), options.WithDefaults());

        /// <summary>
        /// Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
        /// </summary>
        public static Output<Rfc3339tounixResult> Invoke(Rfc3339tounixInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Rfc3339tounixResult>("std:index:rfc3339tounix", args ?? new Rfc3339tounixInvokeArgs(), options.WithDefaults());
    }


    public sealed class Rfc3339tounixArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Rfc3339tounixArgs()
        {
        }
        public static new Rfc3339tounixArgs Empty => new Rfc3339tounixArgs();
    }

    public sealed class Rfc3339tounixInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Rfc3339tounixInvokeArgs()
        {
        }
        public static new Rfc3339tounixInvokeArgs Empty => new Rfc3339tounixInvokeArgs();
    }


    [OutputType]
    public sealed class Rfc3339tounixResult
    {
        public readonly int Result;

        [OutputConstructor]
        private Rfc3339tounixResult(int result)
        {
            Result = result;
        }
    }
}
