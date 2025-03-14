// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the greatest integer value less than or equal to the argument.
func Floor(ctx *pulumi.Context, args *FloorArgs, opts ...pulumi.InvokeOption) (*FloorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FloorResult
	err := ctx.Invoke("std:index:floor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FloorArgs struct {
	Input float64 `pulumi:"input"`
}

type FloorResult struct {
	Result float64 `pulumi:"result"`
}

func FloorOutput(ctx *pulumi.Context, args FloorOutputArgs, opts ...pulumi.InvokeOption) FloorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (FloorResultOutput, error) {
			args := v.(FloorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:floor", args, FloorResultOutput{}, options).(FloorResultOutput), nil
		}).(FloorResultOutput)
}

type FloorOutputArgs struct {
	Input pulumi.Float64Input `pulumi:"input"`
}

func (FloorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FloorArgs)(nil)).Elem()
}

type FloorResultOutput struct{ *pulumi.OutputState }

func (FloorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FloorResult)(nil)).Elem()
}

func (o FloorResultOutput) ToFloorResultOutput() FloorResultOutput {
	return o
}

func (o FloorResultOutput) ToFloorResultOutputWithContext(ctx context.Context) FloorResultOutput {
	return o
}

func (o FloorResultOutput) Result() pulumi.Float64Output {
	return o.ApplyT(func(v FloorResult) float64 { return v.Result }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(FloorResultOutput{})
}
