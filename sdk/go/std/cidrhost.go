// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Takes an IP address range in CIDR notation as input
// and creates an IP address with the given host number.
// If given host number is negative, the count starts from the end of the range.
// For example, cidrhost("10.0.0.0/8", 2) returns 10.0.0.2 and cidrhost("10.0.0.0/8", -2) returns 10.255.255.254.
func Cidrhost(ctx *pulumi.Context, args *CidrhostArgs, opts ...pulumi.InvokeOption) (*CidrhostResult, error) {
	var rv CidrhostResult
	err := ctx.Invoke("std:index:cidrhost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type CidrhostArgs struct {
	Host  int    `pulumi:"host"`
	Input string `pulumi:"input"`
}

type CidrhostResult struct {
	Result string `pulumi:"result"`
}

func CidrhostOutput(ctx *pulumi.Context, args CidrhostOutputArgs, opts ...pulumi.InvokeOption) CidrhostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CidrhostResult, error) {
			args := v.(CidrhostArgs)
			r, err := Cidrhost(ctx, &args, opts...)
			var s CidrhostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CidrhostResultOutput)
}

type CidrhostOutputArgs struct {
	Host  pulumi.IntInput    `pulumi:"host"`
	Input pulumi.StringInput `pulumi:"input"`
}

func (CidrhostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CidrhostArgs)(nil)).Elem()
}

type CidrhostResultOutput struct{ *pulumi.OutputState }

func (CidrhostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CidrhostResult)(nil)).Elem()
}

func (o CidrhostResultOutput) ToCidrhostResultOutput() CidrhostResultOutput {
	return o
}

func (o CidrhostResultOutput) ToCidrhostResultOutputWithContext(ctx context.Context) CidrhostResultOutput {
	return o
}

func (o CidrhostResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v CidrhostResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CidrhostResultOutput{})
}
