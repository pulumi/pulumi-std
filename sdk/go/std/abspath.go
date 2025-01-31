// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns an absolute representation of the specified path.
// If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.
func Abspath(ctx *pulumi.Context, args *AbspathArgs, opts ...pulumi.InvokeOption) (*AbspathResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AbspathResult
	err := ctx.Invoke("std:index:abspath", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type AbspathArgs struct {
	Input string `pulumi:"input"`
}

type AbspathResult struct {
	Result string `pulumi:"result"`
}

func AbspathOutput(ctx *pulumi.Context, args AbspathOutputArgs, opts ...pulumi.InvokeOption) AbspathResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (AbspathResultOutput, error) {
			args := v.(AbspathArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:abspath", args, AbspathResultOutput{}, options).(AbspathResultOutput), nil
		}).(AbspathResultOutput)
}

type AbspathOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (AbspathOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AbspathArgs)(nil)).Elem()
}

type AbspathResultOutput struct{ *pulumi.OutputState }

func (AbspathResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AbspathResult)(nil)).Elem()
}

func (o AbspathResultOutput) ToAbspathResultOutput() AbspathResultOutput {
	return o
}

func (o AbspathResultOutput) ToAbspathResultOutputWithContext(ctx context.Context) AbspathResultOutput {
	return o
}

func (o AbspathResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v AbspathResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AbspathResultOutput{})
}
