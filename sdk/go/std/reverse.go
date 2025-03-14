// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a sequence with the same elements but in reverse order.
func Reverse(ctx *pulumi.Context, args *ReverseArgs, opts ...pulumi.InvokeOption) (*ReverseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ReverseResult
	err := ctx.Invoke("std:index:reverse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ReverseArgs struct {
	Input []interface{} `pulumi:"input"`
}

type ReverseResult struct {
	Result []interface{} `pulumi:"result"`
}

func ReverseOutput(ctx *pulumi.Context, args ReverseOutputArgs, opts ...pulumi.InvokeOption) ReverseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ReverseResultOutput, error) {
			args := v.(ReverseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:reverse", args, ReverseResultOutput{}, options).(ReverseResultOutput), nil
		}).(ReverseResultOutput)
}

type ReverseOutputArgs struct {
	Input pulumi.ArrayInput `pulumi:"input"`
}

func (ReverseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseArgs)(nil)).Elem()
}

type ReverseResultOutput struct{ *pulumi.OutputState }

func (ReverseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReverseResult)(nil)).Elem()
}

func (o ReverseResultOutput) ToReverseResultOutput() ReverseResultOutput {
	return o
}

func (o ReverseResultOutput) ToReverseResultOutputWithContext(ctx context.Context) ReverseResultOutput {
	return o
}

func (o ReverseResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v ReverseResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ReverseResultOutput{})
}
