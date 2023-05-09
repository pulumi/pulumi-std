// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Applies URL encoding to a given string.
func Urlencode(ctx *pulumi.Context, args *UrlencodeArgs, opts ...pulumi.InvokeOption) (*UrlencodeResult, error) {
	var rv UrlencodeResult
	err := ctx.Invoke("std:index:urlencode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type UrlencodeArgs struct {
	Input string `pulumi:"input"`
}

type UrlencodeResult struct {
	Result string `pulumi:"result"`
}

func UrlencodeOutput(ctx *pulumi.Context, args UrlencodeOutputArgs, opts ...pulumi.InvokeOption) UrlencodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UrlencodeResult, error) {
			args := v.(UrlencodeArgs)
			r, err := Urlencode(ctx, &args, opts...)
			var s UrlencodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(UrlencodeResultOutput)
}

type UrlencodeOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (UrlencodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlencodeArgs)(nil)).Elem()
}

type UrlencodeResultOutput struct{ *pulumi.OutputState }

func (UrlencodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UrlencodeResult)(nil)).Elem()
}

func (o UrlencodeResultOutput) ToUrlencodeResultOutput() UrlencodeResultOutput {
	return o
}

func (o UrlencodeResultOutput) ToUrlencodeResultOutputWithContext(ctx context.Context) UrlencodeResultOutput {
	return o
}

func (o UrlencodeResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v UrlencodeResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UrlencodeResultOutput{})
}
