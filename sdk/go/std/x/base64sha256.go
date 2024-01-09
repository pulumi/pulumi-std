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

// Returns a base64-encoded representation of raw SHA-256 sum of the given string.
// This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
func Base64sha256(ctx *pulumi.Context, args *Base64sha256Args, opts ...pulumi.InvokeOption) (*Base64sha256Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Base64sha256Result
	err := ctx.Invoke("std:index:base64sha256", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Base64sha256Args struct {
	Input string `pulumi:"input"`
}

type Base64sha256Result struct {
	Result string `pulumi:"result"`
}

func Base64sha256Output(ctx *pulumi.Context, args Base64sha256OutputArgs, opts ...pulumi.InvokeOption) Base64sha256ResultOutput {
	outputResult := pulumix.ApplyErr[*Base64sha256Args](args.ToOutput(), func(plainArgs *Base64sha256Args) (*Base64sha256Result, error) {
		return Base64sha256(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[Base64sha256ResultOutput, *Base64sha256Result](outputResult)
}

type Base64sha256OutputArgs struct {
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args Base64sha256OutputArgs) ToOutput() pulumix.Output[*Base64sha256Args] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *Base64sha256Args {
		return &Base64sha256Args{
			Input: resolvedArgs[0].(string),
		}
	})
}

type Base64sha256ResultOutput struct{ *pulumi.OutputState }

func (Base64sha256ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Base64sha256Result)(nil)).Elem()
}

func (o Base64sha256ResultOutput) ToOutput(context.Context) pulumix.Output[*Base64sha256Result] {
	return pulumix.Output[*Base64sha256Result]{
		OutputState: o.OutputState,
	}
}

func (o Base64sha256ResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*Base64sha256Result](o, func(v *Base64sha256Result) string { return v.Result })
}