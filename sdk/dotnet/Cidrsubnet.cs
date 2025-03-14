// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Cidrsubnet
    {
        /// <summary>
        /// Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
        /// to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8)
        /// returns 10.2.0.0/16; cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns
        /// 2607:f298:6051:516c:200::/72.
        /// </summary>
        public static Task<CidrsubnetResult> InvokeAsync(CidrsubnetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CidrsubnetResult>("std:index:cidrsubnet", args ?? new CidrsubnetArgs(), options.WithDefaults());

        /// <summary>
        /// Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
        /// to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8)
        /// returns 10.2.0.0/16; cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns
        /// 2607:f298:6051:516c:200::/72.
        /// </summary>
        public static Output<CidrsubnetResult> Invoke(CidrsubnetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CidrsubnetResult>("std:index:cidrsubnet", args ?? new CidrsubnetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Takes an IP address range in CIDR notation (like 10.0.0.0/8) and extends its prefix
        /// to include an additional subnet number. For example, cidrsubnet("10.0.0.0/8", netnum: 2, newbits: 8)
        /// returns 10.2.0.0/16; cidrsubnet("2607:f298:6051:516c::/64", netnum: 2, newbits: 8) returns
        /// 2607:f298:6051:516c:200::/72.
        /// </summary>
        public static Output<CidrsubnetResult> Invoke(CidrsubnetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<CidrsubnetResult>("std:index:cidrsubnet", args ?? new CidrsubnetInvokeArgs(), options.WithDefaults());
    }


    public sealed class CidrsubnetArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        [Input("netnum", required: true)]
        public int Netnum { get; set; }

        [Input("newbits", required: true)]
        public int Newbits { get; set; }

        public CidrsubnetArgs()
        {
        }
        public static new CidrsubnetArgs Empty => new CidrsubnetArgs();
    }

    public sealed class CidrsubnetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        [Input("netnum", required: true)]
        public Input<int> Netnum { get; set; } = null!;

        [Input("newbits", required: true)]
        public Input<int> Newbits { get; set; } = null!;

        public CidrsubnetInvokeArgs()
        {
        }
        public static new CidrsubnetInvokeArgs Empty => new CidrsubnetInvokeArgs();
    }


    [OutputType]
    public sealed class CidrsubnetResult
    {
        public readonly string Result;

        [OutputConstructor]
        private CidrsubnetResult(string result)
        {
            Result = result;
        }
    }
}
