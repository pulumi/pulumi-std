// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Tostring
    {
        /// <summary>
        /// Converts its argument to a string value. Only primitive types (string, number, bool)
        /// 	and null can be converted to string. All other values will result in an error.
        /// </summary>
        public static Task<TostringResult> InvokeAsync(TostringArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<TostringResult>("std:index:tostring", args ?? new TostringArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a string value. Only primitive types (string, number, bool)
        /// 	and null can be converted to string. All other values will result in an error.
        /// </summary>
        public static Output<TostringResult> Invoke(TostringInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<TostringResult>("std:index:tostring", args ?? new TostringInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Converts its argument to a string value. Only primitive types (string, number, bool)
        /// 	and null can be converted to string. All other values will result in an error.
        /// </summary>
        public static Output<TostringResult> Invoke(TostringInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<TostringResult>("std:index:tostring", args ?? new TostringInvokeArgs(), options.WithDefaults());
    }


    public sealed class TostringArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public object Input { get; set; } = null!;

        public TostringArgs()
        {
        }
        public static new TostringArgs Empty => new TostringArgs();
    }

    public sealed class TostringInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<object> Input { get; set; } = null!;

        public TostringInvokeArgs()
        {
        }
        public static new TostringInvokeArgs Empty => new TostringInvokeArgs();
    }


    [OutputType]
    public sealed class TostringResult
    {
        public readonly string? Result;

        [OutputConstructor]
        private TostringResult(string? result)
        {
            Result = result;
        }
    }
}
