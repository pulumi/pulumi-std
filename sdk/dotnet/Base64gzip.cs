// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Base64gzip
    {
        /// <summary>
        /// Compresses the given string with gzip and then encodes the result to base64.
        /// </summary>
        public static Task<Base64gzipResult> InvokeAsync(Base64gzipArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Base64gzipResult>("std:index:base64gzip", args ?? new Base64gzipArgs(), options.WithDefaults());

        /// <summary>
        /// Compresses the given string with gzip and then encodes the result to base64.
        /// </summary>
        public static Output<Base64gzipResult> Invoke(Base64gzipInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Base64gzipResult>("std:index:base64gzip", args ?? new Base64gzipInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Compresses the given string with gzip and then encodes the result to base64.
        /// </summary>
        public static Output<Base64gzipResult> Invoke(Base64gzipInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Base64gzipResult>("std:index:base64gzip", args ?? new Base64gzipInvokeArgs(), options.WithDefaults());
    }


    public sealed class Base64gzipArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Base64gzipArgs()
        {
        }
        public static new Base64gzipArgs Empty => new Base64gzipArgs();
    }

    public sealed class Base64gzipInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Base64gzipInvokeArgs()
        {
        }
        public static new Base64gzipInvokeArgs Empty => new Base64gzipInvokeArgs();
    }


    [OutputType]
    public sealed class Base64gzipResult
    {
        public readonly string Result;

        [OutputConstructor]
        private Base64gzipResult(string result)
        {
            Result = result;
        }
    }
}
