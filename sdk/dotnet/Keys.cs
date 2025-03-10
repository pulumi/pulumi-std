// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Keys
    {
        /// <summary>
        /// Returns a lexically sorted list of the map keys.
        /// </summary>
        public static Task<KeysResult> InvokeAsync(KeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<KeysResult>("std:index:keys", args ?? new KeysArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a lexically sorted list of the map keys.
        /// </summary>
        public static Output<KeysResult> Invoke(KeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<KeysResult>("std:index:keys", args ?? new KeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a lexically sorted list of the map keys.
        /// </summary>
        public static Output<KeysResult> Invoke(KeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<KeysResult>("std:index:keys", args ?? new KeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class KeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private Dictionary<string, object>? _input;
        public Dictionary<string, object> Input
        {
            get => _input ?? (_input = new Dictionary<string, object>());
            set => _input = value;
        }

        public KeysArgs()
        {
        }
        public static new KeysArgs Empty => new KeysArgs();
    }

    public sealed class KeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputMap<object>? _input;
        public InputMap<object> Input
        {
            get => _input ?? (_input = new InputMap<object>());
            set => _input = value;
        }

        public KeysInvokeArgs()
        {
        }
        public static new KeysInvokeArgs Empty => new KeysInvokeArgs();
    }


    [OutputType]
    public sealed class KeysResult
    {
        public readonly ImmutableArray<string> Result;

        [OutputConstructor]
        private KeysResult(ImmutableArray<string> result)
        {
            Result = result;
        }
    }
}
