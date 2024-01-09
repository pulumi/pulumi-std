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

// Parses the given string as a representation of an integer in the specified base
// and returns the resulting number. The base must be between 2 and 62 inclusive.
//
//	.
func Parseint(ctx *pulumi.Context, args *ParseintArgs, opts ...pulumi.InvokeOption) (*ParseintResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ParseintResult
	err := ctx.Invoke("std:index:parseint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ParseintArgs struct {
	Base  *int   `pulumi:"base"`
	Input string `pulumi:"input"`
}

type ParseintResult struct {
	Result int `pulumi:"result"`
}

func ParseintOutput(ctx *pulumi.Context, args ParseintOutputArgs, opts ...pulumi.InvokeOption) ParseintResultOutput {
	outputResult := pulumix.ApplyErr[*ParseintArgs](args.ToOutput(), func(plainArgs *ParseintArgs) (*ParseintResult, error) {
		return Parseint(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[ParseintResultOutput, *ParseintResult](outputResult)
}

type ParseintOutputArgs struct {
	Base  pulumix.Input[*int]   `pulumi:"base"`
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args ParseintOutputArgs) ToOutput() pulumix.Output[*ParseintArgs] {
	allArgs := pulumix.All(
		args.Base.ToOutput(context.Background()).AsAny(),
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *ParseintArgs {
		return &ParseintArgs{
			Base:  resolvedArgs[0].(*int),
			Input: resolvedArgs[1].(string),
		}
	})
}

type ParseintResultOutput struct{ *pulumi.OutputState }

func (ParseintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParseintResult)(nil)).Elem()
}

func (o ParseintResultOutput) ToOutput(context.Context) pulumix.Output[*ParseintResult] {
	return pulumix.Output[*ParseintResult]{
		OutputState: o.OutputState,
	}
}

func (o ParseintResultOutput) Result() pulumix.Output[int] {
	return pulumix.Apply[*ParseintResult](o, func(v *ParseintResult) int { return v.Result })
}