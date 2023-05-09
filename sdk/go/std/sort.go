// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of strings sorted lexicographically.
func Sort(ctx *pulumi.Context, args *SortArgs, opts ...pulumi.InvokeOption) (*SortResult, error) {
	var rv SortResult
	err := ctx.Invoke("std:index:sort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type SortArgs struct {
	Input []string `pulumi:"input"`
}

type SortResult struct {
	Result []string `pulumi:"result"`
}

func SortOutput(ctx *pulumi.Context, args SortOutputArgs, opts ...pulumi.InvokeOption) SortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SortResult, error) {
			args := v.(SortArgs)
			r, err := Sort(ctx, &args, opts...)
			var s SortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SortResultOutput)
}

type SortOutputArgs struct {
	Input pulumi.StringArrayInput `pulumi:"input"`
}

func (SortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SortArgs)(nil)).Elem()
}

type SortResultOutput struct{ *pulumi.OutputState }

func (SortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SortResult)(nil)).Elem()
}

func (o SortResultOutput) ToSortResultOutput() SortResultOutput {
	return o
}

func (o SortResultOutput) ToSortResultOutputWithContext(ctx context.Context) SortResultOutput {
	return o
}

func (o SortResultOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SortResult) []string { return v.Result }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SortResultOutput{})
}
