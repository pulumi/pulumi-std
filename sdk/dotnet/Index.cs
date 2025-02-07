// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Index
    {
        /// <summary>
        /// Finds the index of a given element in a list.
        /// </summary>
        public static Task<IndexResult> InvokeAsync(IndexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<IndexResult>("std:index:index", args ?? new IndexArgs(), options.WithDefaults());

        /// <summary>
        /// Finds the index of a given element in a list.
        /// </summary>
        public static Output<IndexResult> Invoke(IndexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<IndexResult>("std:index:index", args ?? new IndexInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Finds the index of a given element in a list.
        /// </summary>
        public static Output<IndexResult> Invoke(IndexInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<IndexResult>("std:index:index", args ?? new IndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class IndexArgs : global::Pulumi.InvokeArgs
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

        public IndexArgs()
        {
        }
        public static new IndexArgs Empty => new IndexArgs();
    }

    public sealed class IndexInvokeArgs : global::Pulumi.InvokeArgs
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

        public IndexInvokeArgs()
        {
        }
        public static new IndexInvokeArgs Empty => new IndexInvokeArgs();
    }


    [OutputType]
    public sealed class IndexResult
    {
        public readonly int Result;

        [OutputConstructor]
        private IndexResult(int result)
        {
            Result = result;
        }
    }
}
