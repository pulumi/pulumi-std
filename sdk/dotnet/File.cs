// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class File
    {
        /// <summary>
        /// Reads the contents of a file into the string.
        /// </summary>
        public static Task<FileResult> InvokeAsync(FileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<FileResult>("std:index:file", args ?? new FileArgs(), options.WithDefaults());

        /// <summary>
        /// Reads the contents of a file into the string.
        /// </summary>
        public static Output<FileResult> Invoke(FileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<FileResult>("std:index:file", args ?? new FileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Reads the contents of a file into the string.
        /// </summary>
        public static Output<FileResult> Invoke(FileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<FileResult>("std:index:file", args ?? new FileInvokeArgs(), options.WithDefaults());
    }


    public sealed class FileArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public FileArgs()
        {
        }
        public static new FileArgs Empty => new FileArgs();
    }

    public sealed class FileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public FileInvokeArgs()
        {
        }
        public static new FileInvokeArgs Empty => new FileInvokeArgs();
    }


    [OutputType]
    public sealed class FileResult
    {
        public readonly string Result;

        [OutputConstructor]
        private FileResult(string result)
        {
            Result = result;
        }
    }
}
