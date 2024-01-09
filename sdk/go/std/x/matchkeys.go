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

// For two lists values and keys of equal length,
// returns all elements from values where the corresponding element from keys exists in the searchset list.
func Matchkeys(ctx *pulumi.Context, args *MatchkeysArgs, opts ...pulumi.InvokeOption) (*MatchkeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv MatchkeysResult
	err := ctx.Invoke("std:index:matchkeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MatchkeysArgs struct {
	SearchList []string `pulumi:"searchList"`
	Values     []string `pulumi:"values"`
}

type MatchkeysResult struct {
	Result []string `pulumi:"result"`
}

func MatchkeysOutput(ctx *pulumi.Context, args MatchkeysOutputArgs, opts ...pulumi.InvokeOption) MatchkeysResultOutput {
	outputResult := pulumix.ApplyErr[*MatchkeysArgs](args.ToOutput(), func(plainArgs *MatchkeysArgs) (*MatchkeysResult, error) {
		return Matchkeys(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[MatchkeysResultOutput, *MatchkeysResult](outputResult)
}

type MatchkeysOutputArgs struct {
	SearchList pulumix.Input[[]string] `pulumi:"searchList"`
	Values     pulumix.Input[[]string] `pulumi:"values"`
}

func (args MatchkeysOutputArgs) ToOutput() pulumix.Output[*MatchkeysArgs] {
	allArgs := pulumix.All(
		args.SearchList.ToOutput(context.Background()).AsAny(),
		args.Values.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *MatchkeysArgs {
		return &MatchkeysArgs{
			SearchList: resolvedArgs[0].([]string),
			Values:     resolvedArgs[1].([]string),
		}
	})
}

type MatchkeysResultOutput struct{ *pulumi.OutputState }

func (MatchkeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchkeysResult)(nil)).Elem()
}

func (o MatchkeysResultOutput) ToOutput(context.Context) pulumix.Output[*MatchkeysResult] {
	return pulumix.Output[*MatchkeysResult]{
		OutputState: o.OutputState,
	}
}

func (o MatchkeysResultOutput) Result() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[*MatchkeysResult](o, func(v *MatchkeysResult) []string { return v.Result })
	return pulumix.ArrayOutput[string]{
		OutputState: value.OutputState,
	}
}