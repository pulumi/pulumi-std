// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the element at the specified index.
func Element(ctx *pulumi.Context, args *ElementArgs, opts ...pulumi.InvokeOption) (*ElementResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ElementResult
	err := ctx.Invoke("std:index:element", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ElementArgs struct {
	Index int           `pulumi:"index"`
	Input []interface{} `pulumi:"input"`
}

type ElementResult struct {
	Result interface{} `pulumi:"result"`
}

func ElementOutput(ctx *pulumi.Context, args ElementOutputArgs, opts ...pulumi.InvokeOption) ElementResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ElementResultOutput, error) {
			args := v.(ElementArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:element", args, ElementResultOutput{}, options).(ElementResultOutput), nil
		}).(ElementResultOutput)
}

type ElementOutputArgs struct {
	Index pulumi.IntInput   `pulumi:"index"`
	Input pulumi.ArrayInput `pulumi:"input"`
}

func (ElementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ElementArgs)(nil)).Elem()
}

type ElementResultOutput struct{ *pulumi.OutputState }

func (ElementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElementResult)(nil)).Elem()
}

func (o ElementResultOutput) ToElementResultOutput() ElementResultOutput {
	return o
}

func (o ElementResultOutput) ToElementResultOutputWithContext(ctx context.Context) ElementResultOutput {
	return o
}

func (o ElementResultOutput) Result() pulumi.AnyOutput {
	return o.ApplyT(func(v ElementResult) interface{} { return v.Result }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(ElementResultOutput{})
}
