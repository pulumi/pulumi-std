// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Filesha512
    {
        /// <summary>
        /// Reads the contents of a file into a string and returns the SHA512 hash of it.
        /// </summary>
        public static Task<Filesha512Result> InvokeAsync(Filesha512Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Filesha512Result>("std:index:filesha512", args ?? new Filesha512Args(), options.WithDefaults());

        /// <summary>
        /// Reads the contents of a file into a string and returns the SHA512 hash of it.
        /// </summary>
        public static Output<Filesha512Result> Invoke(Filesha512InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Filesha512Result>("std:index:filesha512", args ?? new Filesha512InvokeArgs(), options.WithDefaults());
    }


    public sealed class Filesha512Args : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Filesha512Args()
        {
        }
        public static new Filesha512Args Empty => new Filesha512Args();
    }

    public sealed class Filesha512InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Filesha512InvokeArgs()
        {
        }
        public static new Filesha512InvokeArgs Empty => new Filesha512InvokeArgs();
    }


    [OutputType]
    public sealed class Filesha512Result
    {
        public readonly string Result;

        [OutputConstructor]
        private Filesha512Result(string result)
        {
            Result = result;
        }
    }
}
