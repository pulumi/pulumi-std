// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Base64sha256
    {
        /// <summary>
        /// Returns a base64-encoded representation of raw SHA-256 sum of the given string.
        /// This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
        /// </summary>
        public static Task<Base64sha256Result> InvokeAsync(Base64sha256Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Base64sha256Result>("std:index:base64sha256", args ?? new Base64sha256Args(), options.WithDefaults());

        /// <summary>
        /// Returns a base64-encoded representation of raw SHA-256 sum of the given string.
        /// This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
        /// </summary>
        public static Output<Base64sha256Result> Invoke(Base64sha256InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Base64sha256Result>("std:index:base64sha256", args ?? new Base64sha256InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a base64-encoded representation of raw SHA-256 sum of the given string.
        /// This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
        /// </summary>
        public static Output<Base64sha256Result> Invoke(Base64sha256InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Base64sha256Result>("std:index:base64sha256", args ?? new Base64sha256InvokeArgs(), options.WithDefaults());
    }


    public sealed class Base64sha256Args : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Base64sha256Args()
        {
        }
        public static new Base64sha256Args Empty => new Base64sha256Args();
    }

    public sealed class Base64sha256InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Base64sha256InvokeArgs()
        {
        }
        public static new Base64sha256InvokeArgs Empty => new Base64sha256InvokeArgs();
    }


    [OutputType]
    public sealed class Base64sha256Result
    {
        public readonly string Result;

        [OutputConstructor]
        private Base64sha256Result(string result)
        {
            Result = result;
        }
    }
}
