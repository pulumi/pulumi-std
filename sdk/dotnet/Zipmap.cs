// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Zipmap
    {
        /// <summary>
        /// Constructs a map from a list of keys and a corresponding list of values.
        /// </summary>
        public static Task<ZipmapResult> InvokeAsync(ZipmapArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ZipmapResult>("std:index:zipmap", args ?? new ZipmapArgs(), options.WithDefaults());

        /// <summary>
        /// Constructs a map from a list of keys and a corresponding list of values.
        /// </summary>
        public static Output<ZipmapResult> Invoke(ZipmapInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ZipmapResult>("std:index:zipmap", args ?? new ZipmapInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Constructs a map from a list of keys and a corresponding list of values.
        /// </summary>
        public static Output<ZipmapResult> Invoke(ZipmapInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ZipmapResult>("std:index:zipmap", args ?? new ZipmapInvokeArgs(), options.WithDefaults());
    }


    public sealed class ZipmapArgs : global::Pulumi.InvokeArgs
    {
        [Input("keys", required: true)]
        private List<string>? _keys;
        public List<string> Keys
        {
            get => _keys ?? (_keys = new List<string>());
            set => _keys = value;
        }

        [Input("values", required: true)]
        private List<object>? _values;
        public List<object> Values
        {
            get => _values ?? (_values = new List<object>());
            set => _values = value;
        }

        public ZipmapArgs()
        {
        }
        public static new ZipmapArgs Empty => new ZipmapArgs();
    }

    public sealed class ZipmapInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("keys", required: true)]
        private InputList<string>? _keys;
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        [Input("values", required: true)]
        private InputList<object>? _values;
        public InputList<object> Values
        {
            get => _values ?? (_values = new InputList<object>());
            set => _values = value;
        }

        public ZipmapInvokeArgs()
        {
        }
        public static new ZipmapInvokeArgs Empty => new ZipmapInvokeArgs();
    }


    [OutputType]
    public sealed class ZipmapResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private ZipmapResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
