// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Contains
    {
        /// <summary>
        /// Returns true if a list contains the given element and returns false otherwise.
        /// </summary>
        public static Task<ContainsResult> InvokeAsync(ContainsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ContainsResult>("std:index:contains", args ?? new ContainsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns true if a list contains the given element and returns false otherwise.
        /// </summary>
        public static Output<ContainsResult> Invoke(ContainsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ContainsResult>("std:index:contains", args ?? new ContainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ContainsArgs : global::Pulumi.InvokeArgs
    {
        [Input("element", required: true)]
        public object Element { get; set; } = null!;

        [Input("input", required: true)]
        private List<object>? _input;
        public List<object> Input
        {
            get => _input ?? (_input = new List<object>());
            set => _input = value;
        }

        public ContainsArgs()
        {
        }
        public static new ContainsArgs Empty => new ContainsArgs();
    }

    public sealed class ContainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("element", required: true)]
        public Input<object> Element { get; set; } = null!;

        [Input("input", required: true)]
        private InputList<object>? _input;
        public InputList<object> Input
        {
            get => _input ?? (_input = new InputList<object>());
            set => _input = value;
        }

        public ContainsInvokeArgs()
        {
        }
        public static new ContainsInvokeArgs Empty => new ContainsInvokeArgs();
    }


    [OutputType]
    public sealed class ContainsResult
    {
        public readonly bool Result;

        [OutputConstructor]
        private ContainsResult(bool result)
        {
            Result = result;
        }
    }
}
