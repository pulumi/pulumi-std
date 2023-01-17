// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Extracts a substring from the given string.
func Substr(ctx *pulumi.Context, args *SubstrArgs, opts ...pulumi.InvokeOption) (*SubstrResult, error) {
	var rv SubstrResult
	err := ctx.Invoke("std:index:substr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type SubstrArgs struct {
	Input  string `pulumi:"input"`
	Length int    `pulumi:"length"`
	Offset int    `pulumi:"offset"`
}

type SubstrResult struct {
	Result string `pulumi:"result"`
}

func SubstrOutput(ctx *pulumi.Context, args SubstrOutputArgs, opts ...pulumi.InvokeOption) SubstrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SubstrResult, error) {
			args := v.(SubstrArgs)
			r, err := Substr(ctx, &args, opts...)
			var s SubstrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SubstrResultOutput)
}

type SubstrOutputArgs struct {
	Input  pulumi.StringInput `pulumi:"input"`
	Length pulumi.IntInput    `pulumi:"length"`
	Offset pulumi.IntInput    `pulumi:"offset"`
}

func (SubstrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubstrArgs)(nil)).Elem()
}

type SubstrResultOutput struct{ *pulumi.OutputState }

func (SubstrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubstrResult)(nil)).Elem()
}

func (o SubstrResultOutput) ToSubstrResultOutput() SubstrResultOutput {
	return o
}

func (o SubstrResultOutput) ToSubstrResultOutputWithContext(ctx context.Context) SubstrResultOutput {
	return o
}

func (o SubstrResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v SubstrResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubstrResultOutput{})
}