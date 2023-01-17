// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Compact
    {
        /// <summary>
        /// Removes empty string elements from a list.
        /// </summary>
        public static Task<CompactResult> InvokeAsync(CompactArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CompactResult>("std:index:compact", args ?? new CompactArgs(), options.WithDefaults());

        /// <summary>
        /// Removes empty string elements from a list.
        /// </summary>
        public static Output<CompactResult> Invoke(CompactInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CompactResult>("std:index:compact", args ?? new CompactInvokeArgs(), options.WithDefaults());
    }


    public sealed class CompactArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<string>? _input;
        public List<string> Input
        {
            get => _input ?? (_input = new List<string>());
            set => _input = value;
        }

        public CompactArgs()
        {
        }
        public static new CompactArgs Empty => new CompactArgs();
    }

    public sealed class CompactInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<string>? _input;
        public InputList<string> Input
        {
            get => _input ?? (_input = new InputList<string>());
            set => _input = value;
        }

        public CompactInvokeArgs()
        {
        }
        public static new CompactInvokeArgs Empty => new CompactInvokeArgs();
    }


    [OutputType]
    public sealed class CompactResult
    {
        public readonly ImmutableArray<string> Result;

        [OutputConstructor]
        private CompactResult(ImmutableArray<string> result)
        {
            Result = result;
        }
    }
}
