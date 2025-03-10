// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Determines whether a file exists at a given path.
func Fileexists(ctx *pulumi.Context, args *FileexistsArgs, opts ...pulumi.InvokeOption) (*FileexistsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FileexistsResult
	err := ctx.Invoke("std:index:fileexists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FileexistsArgs struct {
	Input string `pulumi:"input"`
}

type FileexistsResult struct {
	Result bool `pulumi:"result"`
}

func FileexistsOutput(ctx *pulumi.Context, args FileexistsOutputArgs, opts ...pulumi.InvokeOption) FileexistsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (FileexistsResultOutput, error) {
			args := v.(FileexistsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:fileexists", args, FileexistsResultOutput{}, options).(FileexistsResultOutput), nil
		}).(FileexistsResultOutput)
}

type FileexistsOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (FileexistsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileexistsArgs)(nil)).Elem()
}

type FileexistsResultOutput struct{ *pulumi.OutputState }

func (FileexistsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileexistsResult)(nil)).Elem()
}

func (o FileexistsResultOutput) ToFileexistsResultOutput() FileexistsResultOutput {
	return o
}

func (o FileexistsResultOutput) ToFileexistsResultOutputWithContext(ctx context.Context) FileexistsResultOutput {
	return o
}

func (o FileexistsResultOutput) Result() pulumi.BoolOutput {
	return o.ApplyT(func(v FileexistsResult) bool { return v.Result }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(FileexistsResultOutput{})
}
