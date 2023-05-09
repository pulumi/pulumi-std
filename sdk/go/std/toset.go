// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts its argument to a set value.
func Toset(ctx *pulumi.Context, args *TosetArgs, opts ...pulumi.InvokeOption) (*TosetResult, error) {
	var rv TosetResult
	err := ctx.Invoke("std:index:toset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TosetArgs struct {
	Input []interface{} `pulumi:"input"`
}

type TosetResult struct {
	Result []interface{} `pulumi:"result"`
}

func TosetOutput(ctx *pulumi.Context, args TosetOutputArgs, opts ...pulumi.InvokeOption) TosetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TosetResult, error) {
			args := v.(TosetArgs)
			r, err := Toset(ctx, &args, opts...)
			var s TosetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TosetResultOutput)
}

type TosetOutputArgs struct {
	Input pulumi.ArrayInput `pulumi:"input"`
}

func (TosetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TosetArgs)(nil)).Elem()
}

type TosetResultOutput struct{ *pulumi.OutputState }

func (TosetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TosetResult)(nil)).Elem()
}

func (o TosetResultOutput) ToTosetResultOutput() TosetResultOutput {
	return o
}

func (o TosetResultOutput) ToTosetResultOutputWithContext(ctx context.Context) TosetResultOutput {
	return o
}

func (o TosetResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v TosetResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TosetResultOutput{})
}
