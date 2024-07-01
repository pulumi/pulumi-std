// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts a RFC3999 formatted timestamp into a Unix timestamp with milliseconds.
func Unixtorfc3999(ctx *pulumi.Context, args *Unixtorfc3999Args, opts ...pulumi.InvokeOption) (*Unixtorfc3999Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Unixtorfc3999Result
	err := ctx.Invoke("std:index:unixtorfc3999", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Unixtorfc3999Args struct {
	Input string `pulumi:"input"`
}

type Unixtorfc3999Result struct {
	Result int `pulumi:"result"`
}

func Unixtorfc3999Output(ctx *pulumi.Context, args Unixtorfc3999OutputArgs, opts ...pulumi.InvokeOption) Unixtorfc3999ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Unixtorfc3999Result, error) {
			args := v.(Unixtorfc3999Args)
			r, err := Unixtorfc3999(ctx, &args, opts...)
			var s Unixtorfc3999Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Unixtorfc3999ResultOutput)
}

type Unixtorfc3999OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Unixtorfc3999OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Unixtorfc3999Args)(nil)).Elem()
}

type Unixtorfc3999ResultOutput struct{ *pulumi.OutputState }

func (Unixtorfc3999ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Unixtorfc3999Result)(nil)).Elem()
}

func (o Unixtorfc3999ResultOutput) ToUnixtorfc3999ResultOutput() Unixtorfc3999ResultOutput {
	return o
}

func (o Unixtorfc3999ResultOutput) ToUnixtorfc3999ResultOutputWithContext(ctx context.Context) Unixtorfc3999ResultOutput {
	return o
}

func (o Unixtorfc3999ResultOutput) Result() pulumi.IntOutput {
	return o.ApplyT(func(v Unixtorfc3999Result) int { return v.Result }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(Unixtorfc3999ResultOutput{})
}
