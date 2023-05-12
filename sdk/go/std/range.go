// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generates a list of numbers using a start value, a limit value, and a step value.
// Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
// depending on whether limit is greater than or less than start.
func Range(ctx *pulumi.Context, args *RangeArgs, opts ...pulumi.InvokeOption) (*RangeResult, error) {
	var rv RangeResult
	err := ctx.Invoke("std:index:range", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type RangeArgs struct {
	Limit float64  `pulumi:"limit"`
	Start *float64 `pulumi:"start"`
	Step  *float64 `pulumi:"step"`
}

type RangeResult struct {
	Result []float64 `pulumi:"result"`
}

func RangeOutput(ctx *pulumi.Context, args RangeOutputArgs, opts ...pulumi.InvokeOption) RangeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RangeResult, error) {
			args := v.(RangeArgs)
			r, err := Range(ctx, &args, opts...)
			var s RangeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RangeResultOutput)
}

type RangeOutputArgs struct {
	Limit pulumi.Float64Input    `pulumi:"limit"`
	Start pulumi.Float64PtrInput `pulumi:"start"`
	Step  pulumi.Float64PtrInput `pulumi:"step"`
}

func (RangeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RangeArgs)(nil)).Elem()
}

type RangeResultOutput struct{ *pulumi.OutputState }

func (RangeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RangeResult)(nil)).Elem()
}

func (o RangeResultOutput) ToRangeResultOutput() RangeResultOutput {
	return o
}

func (o RangeResultOutput) ToRangeResultOutputWithContext(ctx context.Context) RangeResultOutput {
	return o
}

func (o RangeResultOutput) Result() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v RangeResult) []float64 { return v.Result }).(pulumi.Float64ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(RangeResultOutput{})
}
