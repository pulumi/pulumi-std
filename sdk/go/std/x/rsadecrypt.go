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

// Decrypts an RSA-encrypted ciphertext.
// The cipher text must be base64-encoded and the key must be in PEM format.
func Rsadecrypt(ctx *pulumi.Context, args *RsadecryptArgs, opts ...pulumi.InvokeOption) (*RsadecryptResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RsadecryptResult
	err := ctx.Invoke("std:index:rsadecrypt", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type RsadecryptArgs struct {
	CipherText string `pulumi:"cipherText"`
	Key        string `pulumi:"key"`
}

type RsadecryptResult struct {
	Result string `pulumi:"result"`
}

func RsadecryptOutput(ctx *pulumi.Context, args RsadecryptOutputArgs, opts ...pulumi.InvokeOption) RsadecryptResultOutput {
	outputResult := pulumix.ApplyErr[*RsadecryptArgs](args.ToOutput(), func(plainArgs *RsadecryptArgs) (*RsadecryptResult, error) {
		return Rsadecrypt(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[RsadecryptResultOutput, *RsadecryptResult](outputResult)
}

type RsadecryptOutputArgs struct {
	CipherText pulumix.Input[string] `pulumi:"cipherText"`
	Key        pulumix.Input[string] `pulumi:"key"`
}

func (args RsadecryptOutputArgs) ToOutput() pulumix.Output[*RsadecryptArgs] {
	allArgs := pulumix.All(
		args.CipherText.ToOutput(context.Background()).AsAny(),
		args.Key.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *RsadecryptArgs {
		return &RsadecryptArgs{
			CipherText: resolvedArgs[0].(string),
			Key:        resolvedArgs[1].(string),
		}
	})
}

type RsadecryptResultOutput struct{ *pulumi.OutputState }

func (RsadecryptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RsadecryptResult)(nil)).Elem()
}

func (o RsadecryptResultOutput) ToOutput(context.Context) pulumix.Output[*RsadecryptResult] {
	return pulumix.Output[*RsadecryptResult]{
		OutputState: o.OutputState,
	}
}

func (o RsadecryptResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*RsadecryptResult](o, func(v *RsadecryptResult) string { return v.Result })
}