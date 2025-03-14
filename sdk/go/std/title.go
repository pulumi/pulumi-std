// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts the first letter of each word in the given string to uppercase.
func Title(ctx *pulumi.Context, args *TitleArgs, opts ...pulumi.InvokeOption) (*TitleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TitleResult
	err := ctx.Invoke("std:index:title", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TitleArgs struct {
	Input string `pulumi:"input"`
}

type TitleResult struct {
	Result string `pulumi:"result"`
}

func TitleOutput(ctx *pulumi.Context, args TitleOutputArgs, opts ...pulumi.InvokeOption) TitleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (TitleResultOutput, error) {
			args := v.(TitleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:title", args, TitleResultOutput{}, options).(TitleResultOutput), nil
		}).(TitleResultOutput)
}

type TitleOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (TitleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TitleArgs)(nil)).Elem()
}

type TitleResultOutput struct{ *pulumi.OutputState }

func (TitleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TitleResult)(nil)).Elem()
}

func (o TitleResultOutput) ToTitleResultOutput() TitleResultOutput {
	return o
}

func (o TitleResultOutput) ToTitleResultOutputWithContext(ctx context.Context) TitleResultOutput {
	return o
}

func (o TitleResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TitleResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TitleResultOutput{})
}
