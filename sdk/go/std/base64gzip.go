// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Compresses the given string with gzip and then encodes the result to base64.
func Base64gzip(ctx *pulumi.Context, args *Base64gzipArgs, opts ...pulumi.InvokeOption) (*Base64gzipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Base64gzipResult
	err := ctx.Invoke("std:index:base64gzip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Base64gzipArgs struct {
	Input string `pulumi:"input"`
}

type Base64gzipResult struct {
	Result string `pulumi:"result"`
}

func Base64gzipOutput(ctx *pulumi.Context, args Base64gzipOutputArgs, opts ...pulumi.InvokeOption) Base64gzipResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (Base64gzipResultOutput, error) {
			args := v.(Base64gzipArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:base64gzip", args, Base64gzipResultOutput{}, options).(Base64gzipResultOutput), nil
		}).(Base64gzipResultOutput)
}

type Base64gzipOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Base64gzipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64gzipArgs)(nil)).Elem()
}

type Base64gzipResultOutput struct{ *pulumi.OutputState }

func (Base64gzipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64gzipResult)(nil)).Elem()
}

func (o Base64gzipResultOutput) ToBase64gzipResultOutput() Base64gzipResultOutput {
	return o
}

func (o Base64gzipResultOutput) ToBase64gzipResultOutputWithContext(ctx context.Context) Base64gzipResultOutput {
	return o
}

func (o Base64gzipResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Base64gzipResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Base64gzipResultOutput{})
}
