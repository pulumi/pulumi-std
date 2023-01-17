// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the base input raised to the power of the exponent.
func Pow(ctx *pulumi.Context, args *PowArgs, opts ...pulumi.InvokeOption) (*PowResult, error) {
	var rv PowResult
	err := ctx.Invoke("std:index:pow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type PowArgs struct {
	Base     float64 `pulumi:"base"`
	Exponent float64 `pulumi:"exponent"`
}

type PowResult struct {
	Result float64 `pulumi:"result"`
}

func PowOutput(ctx *pulumi.Context, args PowOutputArgs, opts ...pulumi.InvokeOption) PowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PowResult, error) {
			args := v.(PowArgs)
			r, err := Pow(ctx, &args, opts...)
			var s PowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PowResultOutput)
}

type PowOutputArgs struct {
	Base     pulumi.Float64Input `pulumi:"base"`
	Exponent pulumi.Float64Input `pulumi:"exponent"`
}

func (PowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PowArgs)(nil)).Elem()
}

type PowResultOutput struct{ *pulumi.OutputState }

func (PowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PowResult)(nil)).Elem()
}

func (o PowResultOutput) ToPowResultOutput() PowResultOutput {
	return o
}

func (o PowResultOutput) ToPowResultOutputWithContext(ctx context.Context) PowResultOutput {
	return o
}

func (o PowResultOutput) Result() pulumi.Float64Output {
	return o.ApplyT(func(v PowResult) float64 { return v.Result }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(PowResultOutput{})
}
