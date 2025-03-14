// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Sum
    {
        /// <summary>
        /// Returns the total sum of the elements of the input list.
        /// </summary>
        public static Task<SumResult> InvokeAsync(SumArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SumResult>("std:index:sum", args ?? new SumArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the total sum of the elements of the input list.
        /// </summary>
        public static Output<SumResult> Invoke(SumInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SumResult>("std:index:sum", args ?? new SumInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the total sum of the elements of the input list.
        /// </summary>
        public static Output<SumResult> Invoke(SumInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<SumResult>("std:index:sum", args ?? new SumInvokeArgs(), options.WithDefaults());
    }


    public sealed class SumArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<double>? _input;
        public List<double> Input
        {
            get => _input ?? (_input = new List<double>());
            set => _input = value;
        }

        public SumArgs()
        {
        }
        public static new SumArgs Empty => new SumArgs();
    }

    public sealed class SumInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<double>? _input;
        public InputList<double> Input
        {
            get => _input ?? (_input = new InputList<double>());
            set => _input = value;
        }

        public SumInvokeArgs()
        {
        }
        public static new SumInvokeArgs Empty => new SumInvokeArgs();
    }


    [OutputType]
    public sealed class SumResult
    {
        public readonly double Result;

        [OutputConstructor]
        private SumResult(double result)
        {
            Result = result;
        }
    }
}
