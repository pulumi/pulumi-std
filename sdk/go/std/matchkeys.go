// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// For two lists values and keys of equal length,
// returns all elements from values where the corresponding element from keys exists in the searchset list.
func Matchkeys(ctx *pulumi.Context, args *MatchkeysArgs, opts ...pulumi.InvokeOption) (*MatchkeysResult, error) {
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
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MatchkeysResult, error) {
			args := v.(MatchkeysArgs)
			r, err := Matchkeys(ctx, &args, opts...)
			var s MatchkeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MatchkeysResultOutput)
}

type MatchkeysOutputArgs struct {
	SearchList pulumi.StringArrayInput `pulumi:"searchList"`
	Values     pulumi.StringArrayInput `pulumi:"values"`
}

func (MatchkeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchkeysArgs)(nil)).Elem()
}

type MatchkeysResultOutput struct{ *pulumi.OutputState }

func (MatchkeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchkeysResult)(nil)).Elem()
}

func (o MatchkeysResultOutput) ToMatchkeysResultOutput() MatchkeysResultOutput {
	return o
}

func (o MatchkeysResultOutput) ToMatchkeysResultOutputWithContext(ctx context.Context) MatchkeysResultOutput {
	return o
}

func (o MatchkeysResultOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MatchkeysResult) []string { return v.Result }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(MatchkeysResultOutput{})
}
