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

// Returns a map consisting of the key/value pairs specified as arguments.
func Map(ctx *pulumi.Context, args *MapArgs, opts ...pulumi.InvokeOption) (*MapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv MapResult
	err := ctx.Invoke("std:index:map", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MapArgs struct {
	Args []interface{} `pulumi:"args"`
}

type MapResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func MapOutput(ctx *pulumi.Context, args MapOutputArgs, opts ...pulumi.InvokeOption) MapResultOutput {
	outputResult := pulumix.ApplyErr[*MapArgs](args.ToOutput(), func(plainArgs *MapArgs) (*MapResult, error) {
		return Map(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[MapResultOutput, *MapResult](outputResult)
}

type MapOutputArgs struct {
	Args pulumix.Input[[]interface{}] `pulumi:"args"`
}

func (args MapOutputArgs) ToOutput() pulumix.Output[*MapArgs] {
	allArgs := pulumix.All(
		args.Args.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *MapArgs {
		return &MapArgs{
			Args: resolvedArgs[0].([]interface{}),
		}
	})
}

type MapResultOutput struct{ *pulumi.OutputState }

func (MapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapResult)(nil)).Elem()
}

func (o MapResultOutput) ToOutput(context.Context) pulumix.Output[*MapResult] {
	return pulumix.Output[*MapResult]{
		OutputState: o.OutputState,
	}
}

func (o MapResultOutput) Result() pulumix.MapOutput[interface{}] {
	value := pulumix.Apply[*MapResult](o, func(v *MapResult) map[string]interface{} { return v.Result })
	return pulumix.MapOutput[interface{}]{
		OutputState: value.OutputState,
	}
}
