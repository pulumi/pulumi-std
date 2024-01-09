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

// Returns the given string with all of its Unicode characters in reverse order.
func Strrev(ctx *pulumi.Context, args *StrrevArgs, opts ...pulumi.InvokeOption) (*StrrevResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv StrrevResult
	err := ctx.Invoke("std:index:strrev", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type StrrevArgs struct {
	Input string `pulumi:"input"`
}

type StrrevResult struct {
	Result string `pulumi:"result"`
}

func StrrevOutput(ctx *pulumi.Context, args StrrevOutputArgs, opts ...pulumi.InvokeOption) StrrevResultOutput {
	outputResult := pulumix.ApplyErr[*StrrevArgs](args.ToOutput(), func(plainArgs *StrrevArgs) (*StrrevResult, error) {
		return Strrev(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[StrrevResultOutput, *StrrevResult](outputResult)
}

type StrrevOutputArgs struct {
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args StrrevOutputArgs) ToOutput() pulumix.Output[*StrrevArgs] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *StrrevArgs {
		return &StrrevArgs{
			Input: resolvedArgs[0].(string),
		}
	})
}

type StrrevResultOutput struct{ *pulumi.OutputState }

func (StrrevResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StrrevResult)(nil)).Elem()
}

func (o StrrevResultOutput) ToOutput(context.Context) pulumix.Output[*StrrevResult] {
	return pulumix.Output[*StrrevResult]{
		OutputState: o.OutputState,
	}
}

func (o StrrevResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*StrrevResult](o, func(v *StrrevResult) string { return v.Result })
}