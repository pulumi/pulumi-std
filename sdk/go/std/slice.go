// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the portion of list between from (inclusive) and to (exclusive).
func Slice(ctx *pulumi.Context, args *SliceArgs, opts ...pulumi.InvokeOption) (*SliceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv SliceResult
	err := ctx.Invoke("std:index:slice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type SliceArgs struct {
	From *int          `pulumi:"from"`
	List []interface{} `pulumi:"list"`
	To   *int          `pulumi:"to"`
}

type SliceResult struct {
	Result []interface{} `pulumi:"result"`
}

func SliceOutput(ctx *pulumi.Context, args SliceOutputArgs, opts ...pulumi.InvokeOption) SliceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (SliceResultOutput, error) {
			args := v.(SliceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:slice", args, SliceResultOutput{}, options).(SliceResultOutput), nil
		}).(SliceResultOutput)
}

type SliceOutputArgs struct {
	From pulumi.IntPtrInput `pulumi:"from"`
	List pulumi.ArrayInput  `pulumi:"list"`
	To   pulumi.IntPtrInput `pulumi:"to"`
}

func (SliceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceArgs)(nil)).Elem()
}

type SliceResultOutput struct{ *pulumi.OutputState }

func (SliceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SliceResult)(nil)).Elem()
}

func (o SliceResultOutput) ToSliceResultOutput() SliceResultOutput {
	return o
}

func (o SliceResultOutput) ToSliceResultOutputWithContext(ctx context.Context) SliceResultOutput {
	return o
}

func (o SliceResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v SliceResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SliceResultOutput{})
}
