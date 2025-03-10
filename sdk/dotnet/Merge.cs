// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Merge
    {
        /// <summary>
        /// Returns the union of 2 or more maps. The maps are consumed in the order provided,
        /// and duplicate keys overwrite previous entries.
        /// </summary>
        public static Task<MergeResult> InvokeAsync(MergeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<MergeResult>("std:index:merge", args ?? new MergeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the union of 2 or more maps. The maps are consumed in the order provided,
        /// and duplicate keys overwrite previous entries.
        /// </summary>
        public static Output<MergeResult> Invoke(MergeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<MergeResult>("std:index:merge", args ?? new MergeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the union of 2 or more maps. The maps are consumed in the order provided,
        /// and duplicate keys overwrite previous entries.
        /// </summary>
        public static Output<MergeResult> Invoke(MergeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<MergeResult>("std:index:merge", args ?? new MergeInvokeArgs(), options.WithDefaults());
    }


    public sealed class MergeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<ImmutableDictionary<string, object>>? _input;
        public List<ImmutableDictionary<string, object>> Input
        {
            get => _input ?? (_input = new List<ImmutableDictionary<string, object>>());
            set => _input = value;
        }

        public MergeArgs()
        {
        }
        public static new MergeArgs Empty => new MergeArgs();
    }

    public sealed class MergeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<ImmutableDictionary<string, object>>? _input;
        public InputList<ImmutableDictionary<string, object>> Input
        {
            get => _input ?? (_input = new InputList<ImmutableDictionary<string, object>>());
            set => _input = value;
        }

        public MergeInvokeArgs()
        {
        }
        public static new MergeInvokeArgs Empty => new MergeInvokeArgs();
    }


    [OutputType]
    public sealed class MergeResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private MergeResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
