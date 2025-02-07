// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Pow
    {
        /// <summary>
        /// Returns the base input raised to the power of the exponent.
        /// </summary>
        public static Task<PowResult> InvokeAsync(PowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<PowResult>("std:index:pow", args ?? new PowArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the base input raised to the power of the exponent.
        /// </summary>
        public static Output<PowResult> Invoke(PowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<PowResult>("std:index:pow", args ?? new PowInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the base input raised to the power of the exponent.
        /// </summary>
        public static Output<PowResult> Invoke(PowInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<PowResult>("std:index:pow", args ?? new PowInvokeArgs(), options.WithDefaults());
    }


    public sealed class PowArgs : global::Pulumi.InvokeArgs
    {
        [Input("base", required: true)]
        public double Base { get; set; }

        [Input("exponent", required: true)]
        public double Exponent { get; set; }

        public PowArgs()
        {
        }
        public static new PowArgs Empty => new PowArgs();
    }

    public sealed class PowInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("base", required: true)]
        public Input<double> Base { get; set; } = null!;

        [Input("exponent", required: true)]
        public Input<double> Exponent { get; set; } = null!;

        public PowInvokeArgs()
        {
        }
        public static new PowInvokeArgs Empty => new PowInvokeArgs();
    }


    [OutputType]
    public sealed class PowResult
    {
        public readonly double Result;

        [OutputConstructor]
        private PowResult(double result)
        {
            Result = result;
        }
    }
}
