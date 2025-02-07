// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a hexadecimal representation of the SHA-512 hash of the given string.
func Sha512(ctx *pulumi.Context, args *Sha512Args, opts ...pulumi.InvokeOption) (*Sha512Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Sha512Result
	err := ctx.Invoke("std:index:sha512", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Sha512Args struct {
	Input string `pulumi:"input"`
}

type Sha512Result struct {
	Result string `pulumi:"result"`
}

func Sha512Output(ctx *pulumi.Context, args Sha512OutputArgs, opts ...pulumi.InvokeOption) Sha512ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (Sha512ResultOutput, error) {
			args := v.(Sha512Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("std:index:sha512", args, Sha512ResultOutput{}, options).(Sha512ResultOutput), nil
		}).(Sha512ResultOutput)
}

type Sha512OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Sha512OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sha512Args)(nil)).Elem()
}

type Sha512ResultOutput struct{ *pulumi.OutputState }

func (Sha512ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sha512Result)(nil)).Elem()
}

func (o Sha512ResultOutput) ToSha512ResultOutput() Sha512ResultOutput {
	return o
}

func (o Sha512ResultOutput) ToSha512ResultOutputWithContext(ctx context.Context) Sha512ResultOutput {
	return o
}

func (o Sha512ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Sha512Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Sha512ResultOutput{})
}
