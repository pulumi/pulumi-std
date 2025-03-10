// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Joins the list with the delimiter for a resultant string.
func Join(ctx *pulumi.Context, args *JoinArgs, opts ...pulumi.InvokeOption) (*JoinResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv JoinResult
	err := ctx.Invoke("std:index:join", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type JoinArgs struct {
	Input     []string `pulumi:"input"`
	Separator string   `pulumi:"separator"`
}

type JoinResult struct {
	Result string `pulumi:"result"`
}

func JoinOutput(ctx *pulumi.Context, args JoinOutputArgs, opts ...pulumi.InvokeOption) JoinResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (JoinResultOutput, error) {
			args := v.(JoinArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:join", args, JoinResultOutput{}, options).(JoinResultOutput), nil
		}).(JoinResultOutput)
}

type JoinOutputArgs struct {
	Input     pulumi.StringArrayInput `pulumi:"input"`
	Separator pulumi.StringInput      `pulumi:"separator"`
}

func (JoinOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JoinArgs)(nil)).Elem()
}

type JoinResultOutput struct{ *pulumi.OutputState }

func (JoinResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JoinResult)(nil)).Elem()
}

func (o JoinResultOutput) ToJoinResultOutput() JoinResultOutput {
	return o
}

func (o JoinResultOutput) ToJoinResultOutputWithContext(ctx context.Context) JoinResultOutput {
	return o
}

func (o JoinResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v JoinResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JoinResultOutput{})
}
