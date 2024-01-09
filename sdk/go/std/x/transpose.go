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

// Takes a map of lists of strings and swaps the keys and values to return a new map of lists of strings.
func Transpose(ctx *pulumi.Context, args *TransposeArgs, opts ...pulumi.InvokeOption) (*TransposeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TransposeResult
	err := ctx.Invoke("std:index:transpose", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TransposeArgs struct {
	Input map[string][]string `pulumi:"input"`
}

type TransposeResult struct {
	Result map[string][]string `pulumi:"result"`
}

func TransposeOutput(ctx *pulumi.Context, args TransposeOutputArgs, opts ...pulumi.InvokeOption) TransposeResultOutput {
	outputResult := pulumix.ApplyErr[*TransposeArgs](args.ToOutput(), func(plainArgs *TransposeArgs) (*TransposeResult, error) {
		return Transpose(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[TransposeResultOutput, *TransposeResult](outputResult)
}

type TransposeOutputArgs struct {
	Input pulumix.Input[map[string][]string] `pulumi:"input"`
}

func (args TransposeOutputArgs) ToOutput() pulumix.Output[*TransposeArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *TransposeArgs {
		return &TransposeArgs{
			Input: resolvedArgs[0].(map[string][]string),
		}
	})
}

type TransposeResultOutput struct{ *pulumi.OutputState }

func (TransposeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransposeResult)(nil)).Elem()
}

func (o TransposeResultOutput) ToOutput(context.Context) pulumix.Output[*TransposeResult] {
	return pulumix.Output[*TransposeResult]{
		OutputState: o.OutputState,
	}
}

func (o TransposeResultOutput) Result() pulumix.GMapOutput[[]string, []stringOutput] {
	value := pulumix.Apply[*TransposeResult](o, func(v *TransposeResult) map[string][]string { return v.Result })
	return pulumix.GMapOutput[[]string, []stringOutput]{
		OutputState: value.OutputState,
	}
}