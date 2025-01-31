// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Tonumber
    {
        /// <summary>
        /// Converts its argument to a number value. Only number values, null, and strings
        /// 	containing decimal representations of numbers can be converted to number. All other values will result in an error
        /// </summary>
        public static Task<TonumberResult> InvokeAsync(TonumberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<TonumberResult>("std:index:tonumber", args ?? new TonumberArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a number value. Only number values, null, and strings
        /// 	containing decimal representations of numbers can be converted to number. All other values will result in an error
        /// </summary>
        public static Output<TonumberResult> Invoke(TonumberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<TonumberResult>("std:index:tonumber", args ?? new TonumberInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a number value. Only number values, null, and strings
        /// 	containing decimal representations of numbers can be converted to number. All other values will result in an error
        /// </summary>
        public static Output<TonumberResult> Invoke(TonumberInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<TonumberResult>("std:index:tonumber", args ?? new TonumberInvokeArgs(), options.WithDefaults());
    }


    public sealed class TonumberArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public object Input { get; set; } = null!;

        public TonumberArgs()
        {
        }
        public static new TonumberArgs Empty => new TonumberArgs();
    }

    public sealed class TonumberInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<object> Input { get; set; } = null!;

        public TonumberInvokeArgs()
        {
        }
        public static new TonumberInvokeArgs Empty => new TonumberInvokeArgs();
    }


    [OutputType]
    public sealed class TonumberResult
    {
        public readonly double? Result;

        [OutputConstructor]
        private TonumberResult(double? result)
        {
            Result = result;
        }
    }
}
