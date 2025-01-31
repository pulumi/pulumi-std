// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a base64-encoded representation of raw SHA-256 sum of the given string.
// This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
func Base64sha256(ctx *pulumi.Context, args *Base64sha256Args, opts ...pulumi.InvokeOption) (*Base64sha256Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Base64sha256Result
	err := ctx.Invoke("std:index:base64sha256", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Base64sha256Args struct {
	Input string `pulumi:"input"`
}

type Base64sha256Result struct {
	Result string `pulumi:"result"`
}

func Base64sha256Output(ctx *pulumi.Context, args Base64sha256OutputArgs, opts ...pulumi.InvokeOption) Base64sha256ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (Base64sha256ResultOutput, error) {
			args := v.(Base64sha256Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:base64sha256", args, Base64sha256ResultOutput{}, options).(Base64sha256ResultOutput), nil
		}).(Base64sha256ResultOutput)
}

type Base64sha256OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Base64sha256OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64sha256Args)(nil)).Elem()
}

type Base64sha256ResultOutput struct{ *pulumi.OutputState }

func (Base64sha256ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64sha256Result)(nil)).Elem()
}

func (o Base64sha256ResultOutput) ToBase64sha256ResultOutput() Base64sha256ResultOutput {
	return o
}

func (o Base64sha256ResultOutput) ToBase64sha256ResultOutputWithContext(ctx context.Context) Base64sha256ResultOutput {
	return o
}

func (o Base64sha256ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Base64sha256Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Base64sha256ResultOutput{})
}
