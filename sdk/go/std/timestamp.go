// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a UTC timestamp string of the current time in RFC 3339 format
func Timestamp(ctx *pulumi.Context, args *TimestampArgs, opts ...pulumi.InvokeOption) (*TimestampResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TimestampResult
	err := ctx.Invoke("std:index:timestamp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TimestampArgs struct {
}

type TimestampResult struct {
	Result string `pulumi:"result"`
}

func TimestampOutput(ctx *pulumi.Context, args TimestampOutputArgs, opts ...pulumi.InvokeOption) TimestampResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (TimestampResultOutput, error) {
			args := v.(TimestampArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:timestamp", args, TimestampResultOutput{}, options).(TimestampResultOutput), nil
		}).(TimestampResultOutput)
}

type TimestampOutputArgs struct {
}

func (TimestampOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimestampArgs)(nil)).Elem()
}

type TimestampResultOutput struct{ *pulumi.OutputState }

func (TimestampResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimestampResult)(nil)).Elem()
}

func (o TimestampResultOutput) ToTimestampResultOutput() TimestampResultOutput {
	return o
}

func (o TimestampResultOutput) ToTimestampResultOutputWithContext(ctx context.Context) TimestampResultOutput {
	return o
}

func (o TimestampResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TimestampResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TimestampResultOutput{})
}
