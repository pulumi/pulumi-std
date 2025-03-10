// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Interprets a given string as JSON and returns a represetation
//
//	of the result of decoding that string.
//	If input is not valid JSON, the result will be the input unchanged.
func Jsondecode(ctx *pulumi.Context, args *JsondecodeArgs, opts ...pulumi.InvokeOption) (*JsondecodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv JsondecodeResult
	err := ctx.Invoke("std:index:jsondecode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type JsondecodeArgs struct {
	Input string `pulumi:"input"`
}

type JsondecodeResult struct {
	Result interface{} `pulumi:"result"`
}

func JsondecodeOutput(ctx *pulumi.Context, args JsondecodeOutputArgs, opts ...pulumi.InvokeOption) JsondecodeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (JsondecodeResultOutput, error) {
			args := v.(JsondecodeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:jsondecode", args, JsondecodeResultOutput{}, options).(JsondecodeResultOutput), nil
		}).(JsondecodeResultOutput)
}

type JsondecodeOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (JsondecodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsondecodeArgs)(nil)).Elem()
}

type JsondecodeResultOutput struct{ *pulumi.OutputState }

func (JsondecodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsondecodeResult)(nil)).Elem()
}

func (o JsondecodeResultOutput) ToJsondecodeResultOutput() JsondecodeResultOutput {
	return o
}

func (o JsondecodeResultOutput) ToJsondecodeResultOutputWithContext(ctx context.Context) JsondecodeResultOutput {
	return o
}

func (o JsondecodeResultOutput) Result() pulumi.AnyOutput {
	return o.ApplyT(func(v JsondecodeResult) interface{} { return v.Result }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(JsondecodeResultOutput{})
}
