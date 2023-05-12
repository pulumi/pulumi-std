// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the last element of the input path.
func Basename(ctx *pulumi.Context, args *BasenameArgs, opts ...pulumi.InvokeOption) (*BasenameResult, error) {
	var rv BasenameResult
	err := ctx.Invoke("std:index:basename", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type BasenameArgs struct {
	Input string `pulumi:"input"`
}

type BasenameResult struct {
	Result string `pulumi:"result"`
}

func BasenameOutput(ctx *pulumi.Context, args BasenameOutputArgs, opts ...pulumi.InvokeOption) BasenameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BasenameResult, error) {
			args := v.(BasenameArgs)
			r, err := Basename(ctx, &args, opts...)
			var s BasenameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BasenameResultOutput)
}

type BasenameOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (BasenameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BasenameArgs)(nil)).Elem()
}

type BasenameResultOutput struct{ *pulumi.OutputState }

func (BasenameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasenameResult)(nil)).Elem()
}

func (o BasenameResultOutput) ToBasenameResultOutput() BasenameResultOutput {
	return o
}

func (o BasenameResultOutput) ToBasenameResultOutputWithContext(ctx context.Context) BasenameResultOutput {
	return o
}

func (o BasenameResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v BasenameResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BasenameResultOutput{})
}
