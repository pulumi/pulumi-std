// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the absolute value of a given float.
// Example: abs(1) returns 1, and abs(-1) would also return 1, whereas abs(-3.14) would return 3.14.
func Abs(ctx *pulumi.Context, args *AbsArgs, opts ...pulumi.InvokeOption) (*AbsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AbsResult
	err := ctx.Invoke("std:index:abs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type AbsArgs struct {
	Input float64 `pulumi:"input"`
}

type AbsResult struct {
	Result float64 `pulumi:"result"`
}

func AbsOutput(ctx *pulumi.Context, args AbsOutputArgs, opts ...pulumi.InvokeOption) AbsResultOutput {
	outputResult := pulumix.ApplyErr[*AbsArgs](args.ToOutput(), func(plainArgs *AbsArgs) (*AbsResult, error) {
		return Abs(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[AbsResultOutput, *AbsResult](outputResult)
}

type AbsOutputArgs struct {
	Input pulumix.Input[float64] `pulumi:"input"`
}

func (args AbsOutputArgs) ToOutput() pulumix.Output[*AbsArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *AbsArgs {
		return &AbsArgs{
			Input: resolvedArgs[0].(float64),
		}
	})
}

type AbsResultOutput struct{ *pulumi.OutputState }

func (AbsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsResult)(nil)).Elem()
}

func (o AbsResultOutput) ToOutput(context.Context) pulumix.Output[*AbsResult] {
	return pulumix.Output[*AbsResult]{
		OutputState: o.OutputState,
	}
}

func (o AbsResultOutput) Result() pulumix.Output[float64] {
	return pulumix.Apply[*AbsResult](o, func(v *AbsResult) float64 { return v.Result })
}