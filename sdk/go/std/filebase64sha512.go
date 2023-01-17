// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the contents of a file into a string and returns the base64-encoded SHA512 hash of it.
func Filebase64sha512(ctx *pulumi.Context, args *Filebase64sha512Args, opts ...pulumi.InvokeOption) (*Filebase64sha512Result, error) {
	var rv Filebase64sha512Result
	err := ctx.Invoke("std:index:filebase64sha512", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filebase64sha512Args struct {
	Input string `pulumi:"input"`
}

type Filebase64sha512Result struct {
	Result string `pulumi:"result"`
}

func Filebase64sha512Output(ctx *pulumi.Context, args Filebase64sha512OutputArgs, opts ...pulumi.InvokeOption) Filebase64sha512ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Filebase64sha512Result, error) {
			args := v.(Filebase64sha512Args)
			r, err := Filebase64sha512(ctx, &args, opts...)
			var s Filebase64sha512Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Filebase64sha512ResultOutput)
}

type Filebase64sha512OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Filebase64sha512OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filebase64sha512Args)(nil)).Elem()
}

type Filebase64sha512ResultOutput struct{ *pulumi.OutputState }

func (Filebase64sha512ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filebase64sha512Result)(nil)).Elem()
}

func (o Filebase64sha512ResultOutput) ToFilebase64sha512ResultOutput() Filebase64sha512ResultOutput {
	return o
}

func (o Filebase64sha512ResultOutput) ToFilebase64sha512ResultOutputWithContext(ctx context.Context) Filebase64sha512ResultOutput {
	return o
}

func (o Filebase64sha512ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Filebase64sha512Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Filebase64sha512ResultOutput{})
}
