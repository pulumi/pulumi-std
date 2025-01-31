// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Tolist
    {
        /// <summary>
        /// Converts its argument to a list value.
        /// </summary>
        public static Task<TolistResult> InvokeAsync(TolistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<TolistResult>("std:index:tolist", args ?? new TolistArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a list value.
        /// </summary>
        public static Output<TolistResult> Invoke(TolistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<TolistResult>("std:index:tolist", args ?? new TolistInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a list value.
        /// </summary>
        public static Output<TolistResult> Invoke(TolistInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<TolistResult>("std:index:tolist", args ?? new TolistInvokeArgs(), options.WithDefaults());
    }


    public sealed class TolistArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<object>? _input;
        public List<object> Input
        {
            get => _input ?? (_input = new List<object>());
            set => _input = value;
        }

        public TolistArgs()
        {
        }
        public static new TolistArgs Empty => new TolistArgs();
    }

    public sealed class TolistInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<object>? _input;
        public InputList<object> Input
        {
            get => _input ?? (_input = new InputList<object>());
            set => _input = value;
        }

        public TolistInvokeArgs()
        {
        }
        public static new TolistInvokeArgs Empty => new TolistInvokeArgs();
    }


    [OutputType]
    public sealed class TolistResult
    {
        public readonly ImmutableArray<object> Result;

        [OutputConstructor]
        private TolistResult(ImmutableArray<object> result)
        {
            Result = result;
        }
    }
}
