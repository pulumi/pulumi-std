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

// Reads the contents of a file into a string and returns the MD5 hash of it.
func Filemd5(ctx *pulumi.Context, args *Filemd5Args, opts ...pulumi.InvokeOption) (*Filemd5Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Filemd5Result
	err := ctx.Invoke("std:index:filemd5", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filemd5Args struct {
	Input string `pulumi:"input"`
}

type Filemd5Result struct {
	Result string `pulumi:"result"`
}

func Filemd5Output(ctx *pulumi.Context, args Filemd5OutputArgs, opts ...pulumi.InvokeOption) Filemd5ResultOutput {
	outputResult := pulumix.ApplyErr[*Filemd5Args](args.ToOutput(), func(plainArgs *Filemd5Args) (*Filemd5Result, error) {
		return Filemd5(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[Filemd5ResultOutput, *Filemd5Result](outputResult)
}

type Filemd5OutputArgs struct {
	Input pulumix.Input[string] `pulumi:"input"`
}

func (args Filemd5OutputArgs) ToOutput() pulumix.Output[*Filemd5Args] {
	allArgs := pulumix.All(
		args.Input.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *Filemd5Args {
		return &Filemd5Args{
			Input: resolvedArgs[0].(string),
		}
	})
}

type Filemd5ResultOutput struct{ *pulumi.OutputState }

func (Filemd5ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filemd5Result)(nil)).Elem()
}

func (o Filemd5ResultOutput) ToOutput(context.Context) pulumix.Output[*Filemd5Result] {
	return pulumix.Output[*Filemd5Result]{
		OutputState: o.OutputState,
	}
}

func (o Filemd5ResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*Filemd5Result](o, func(v *Filemd5Result) string { return v.Result })
}