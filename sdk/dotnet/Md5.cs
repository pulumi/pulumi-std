// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Md5
    {
        /// <summary>
        /// Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.
        /// </summary>
        public static Task<Md5Result> InvokeAsync(Md5Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Md5Result>("std:index:md5", args ?? new Md5Args(), options.WithDefaults());

        /// <summary>
        /// Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.
        /// </summary>
        public static Output<Md5Result> Invoke(Md5InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Md5Result>("std:index:md5", args ?? new Md5InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a (conventional) hexadecimal representation of the MD5 hash of the given string.
        /// </summary>
        public static Output<Md5Result> Invoke(Md5InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Md5Result>("std:index:md5", args ?? new Md5InvokeArgs(), options.WithDefaults());
    }


    public sealed class Md5Args : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Md5Args()
        {
        }
        public static new Md5Args Empty => new Md5Args();
    }

    public sealed class Md5InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Md5InvokeArgs()
        {
        }
        public static new Md5InvokeArgs Empty => new Md5InvokeArgs();
    }


    [OutputType]
    public sealed class Md5Result
    {
        public readonly string Result;

        [OutputConstructor]
        private Md5Result(string result)
        {
            Result = result;
        }
    }
}
