// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Jsondecode
    {
        /// <summary>
        /// Interprets a given string as JSON and returns a represetation
        /// 	of the result of decoding that string.
        /// 	If input is not valid JSON, the result will be the input unchanged.
        /// </summary>
        public static Task<JsondecodeResult> InvokeAsync(JsondecodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<JsondecodeResult>("std:index:jsondecode", args ?? new JsondecodeArgs(), options.WithDefaults());

        /// <summary>
        /// Interprets a given string as JSON and returns a represetation
        /// 	of the result of decoding that string.
        /// 	If input is not valid JSON, the result will be the input unchanged.
        /// </summary>
        public static Output<JsondecodeResult> Invoke(JsondecodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<JsondecodeResult>("std:index:jsondecode", args ?? new JsondecodeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Interprets a given string as JSON and returns a represetation
        /// 	of the result of decoding that string.
        /// 	If input is not valid JSON, the result will be the input unchanged.
        /// </summary>
        public static Output<JsondecodeResult> Invoke(JsondecodeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<JsondecodeResult>("std:index:jsondecode", args ?? new JsondecodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class JsondecodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public JsondecodeArgs()
        {
        }
        public static new JsondecodeArgs Empty => new JsondecodeArgs();
    }

    public sealed class JsondecodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public JsondecodeInvokeArgs()
        {
        }
        public static new JsondecodeInvokeArgs Empty => new JsondecodeInvokeArgs();
    }


    [OutputType]
    public sealed class JsondecodeResult
    {
        public readonly object Result;

        [OutputConstructor]
        private JsondecodeResult(object result)
        {
            Result = result;
        }
    }
}
