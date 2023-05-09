// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Combines two or more lists into a single list.
func Concat(ctx *pulumi.Context, args *ConcatArgs, opts ...pulumi.InvokeOption) (*ConcatResult, error) {
	var rv ConcatResult
	err := ctx.Invoke("std:index:concat", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ConcatArgs struct {
	Input [][]interface{} `pulumi:"input"`
}

type ConcatResult struct {
	Result []interface{} `pulumi:"result"`
}

func ConcatOutput(ctx *pulumi.Context, args ConcatOutputArgs, opts ...pulumi.InvokeOption) ConcatResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConcatResult, error) {
			args := v.(ConcatArgs)
			r, err := Concat(ctx, &args, opts...)
			var s ConcatResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConcatResultOutput)
}

type ConcatOutputArgs struct {
	Input pulumi.ArrayArrayInput `pulumi:"input"`
}

func (ConcatOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConcatArgs)(nil)).Elem()
}

type ConcatResultOutput struct{ *pulumi.OutputState }

func (ConcatResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConcatResult)(nil)).Elem()
}

func (o ConcatResultOutput) ToConcatResultOutput() ConcatResultOutput {
	return o
}

func (o ConcatResultOutput) ToConcatResultOutputWithContext(ctx context.Context) ConcatResultOutput {
	return o
}

func (o ConcatResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v ConcatResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ConcatResultOutput{})
}
