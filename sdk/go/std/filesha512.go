// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the contents of a file into a string and returns the SHA512 hash of it.
func Filesha512(ctx *pulumi.Context, args *Filesha512Args, opts ...pulumi.InvokeOption) (*Filesha512Result, error) {
	var rv Filesha512Result
	err := ctx.Invoke("std:index:filesha512", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filesha512Args struct {
	Input string `pulumi:"input"`
}

type Filesha512Result struct {
	Result string `pulumi:"result"`
}

func Filesha512Output(ctx *pulumi.Context, args Filesha512OutputArgs, opts ...pulumi.InvokeOption) Filesha512ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Filesha512Result, error) {
			args := v.(Filesha512Args)
			r, err := Filesha512(ctx, &args, opts...)
			var s Filesha512Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Filesha512ResultOutput)
}

type Filesha512OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Filesha512OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filesha512Args)(nil)).Elem()
}

type Filesha512ResultOutput struct{ *pulumi.OutputState }

func (Filesha512ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filesha512Result)(nil)).Elem()
}

func (o Filesha512ResultOutput) ToFilesha512ResultOutput() Filesha512ResultOutput {
	return o
}

func (o Filesha512ResultOutput) ToFilesha512ResultOutputWithContext(ctx context.Context) Filesha512ResultOutput {
	return o
}

func (o Filesha512ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Filesha512Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Filesha512ResultOutput{})
}
