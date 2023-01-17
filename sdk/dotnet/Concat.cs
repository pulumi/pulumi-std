// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Concat
    {
        /// <summary>
        /// Combines two or more lists into a single list.
        /// </summary>
        public static Task<ConcatResult> InvokeAsync(ConcatArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ConcatResult>("std:index:concat", args ?? new ConcatArgs(), options.WithDefaults());

        /// <summary>
        /// Combines two or more lists into a single list.
        /// </summary>
        public static Output<ConcatResult> Invoke(ConcatInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ConcatResult>("std:index:concat", args ?? new ConcatInvokeArgs(), options.WithDefaults());
    }


    public sealed class ConcatArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<ImmutableArray<object>>? _input;
        public List<ImmutableArray<object>> Input
        {
            get => _input ?? (_input = new List<ImmutableArray<object>>());
            set => _input = value;
        }

        public ConcatArgs()
        {
        }
        public static new ConcatArgs Empty => new ConcatArgs();
    }

    public sealed class ConcatInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<ImmutableArray<object>>? _input;
        public InputList<ImmutableArray<object>> Input
        {
            get => _input ?? (_input = new InputList<ImmutableArray<object>>());
            set => _input = value;
        }

        public ConcatInvokeArgs()
        {
        }
        public static new ConcatInvokeArgs Empty => new ConcatInvokeArgs();
    }


    [OutputType]
    public sealed class ConcatResult
    {
        public readonly ImmutableArray<object> Result;

        [OutputConstructor]
        private ConcatResult(ImmutableArray<object> result)
        {
            Result = result;
        }
    }
}
