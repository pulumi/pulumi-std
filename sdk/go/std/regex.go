// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the first match of a regular expression in a string (including named or indexed groups).
func Regex(ctx *pulumi.Context, args *RegexArgs, opts ...pulumi.InvokeOption) (*RegexResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RegexResult
	err := ctx.Invoke("std:index:regex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type RegexArgs struct {
	Pattern string `pulumi:"pattern"`
	String  string `pulumi:"string"`
}

type RegexResult struct {
	Result interface{} `pulumi:"result"`
}

func RegexOutput(ctx *pulumi.Context, args RegexOutputArgs, opts ...pulumi.InvokeOption) RegexResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (RegexResultOutput, error) {
			args := v.(RegexArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:regex", args, RegexResultOutput{}, options).(RegexResultOutput), nil
		}).(RegexResultOutput)
}

type RegexOutputArgs struct {
	Pattern pulumi.StringInput `pulumi:"pattern"`
	String  pulumi.StringInput `pulumi:"string"`
}

func (RegexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexArgs)(nil)).Elem()
}

type RegexResultOutput struct{ *pulumi.OutputState }

func (RegexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexResult)(nil)).Elem()
}

func (o RegexResultOutput) ToRegexResultOutput() RegexResultOutput {
	return o
}

func (o RegexResultOutput) ToRegexResultOutputWithContext(ctx context.Context) RegexResultOutput {
	return o
}

func (o RegexResultOutput) Result() pulumi.AnyOutput {
	return o.ApplyT(func(v RegexResult) interface{} { return v.Result }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(RegexResultOutput{})
}
