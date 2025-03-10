// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Base64decode
    {
        /// <summary>
        /// Given a base64-encoded string, decodes it and returns the original string.
        /// </summary>
        public static Task<Base64decodeResult> InvokeAsync(Base64decodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Base64decodeResult>("std:index:base64decode", args ?? new Base64decodeArgs(), options.WithDefaults());

        /// <summary>
        /// Given a base64-encoded string, decodes it and returns the original string.
        /// </summary>
        public static Output<Base64decodeResult> Invoke(Base64decodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Base64decodeResult>("std:index:base64decode", args ?? new Base64decodeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Given a base64-encoded string, decodes it and returns the original string.
        /// </summary>
        public static Output<Base64decodeResult> Invoke(Base64decodeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Base64decodeResult>("std:index:base64decode", args ?? new Base64decodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class Base64decodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Base64decodeArgs()
        {
        }
        public static new Base64decodeArgs Empty => new Base64decodeArgs();
    }

    public sealed class Base64decodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Base64decodeInvokeArgs()
        {
        }
        public static new Base64decodeInvokeArgs Empty => new Base64decodeInvokeArgs();
    }


    [OutputType]
    public sealed class Base64decodeResult
    {
        public readonly string Result;

        [OutputConstructor]
        private Base64decodeResult(string result)
        {
            Result = result;
        }
    }
}
