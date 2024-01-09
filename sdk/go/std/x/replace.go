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

// Does a search and replace on the given string.
// All instances of search are replaced with the value of replace.
// If search is wrapped in forward slashes, it is treated as a regular expression.
// If using a regular expression, replace can reference subcaptures in the regular expression by
// using $n where n is the index or name of the subcapture. If using a regular expression,
// the syntax conforms to the re2 regular expression syntax.
func Replace(ctx *pulumi.Context, args *ReplaceArgs, opts ...pulumi.InvokeOption) (*ReplaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ReplaceResult
	err := ctx.Invoke("std:index:replace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ReplaceArgs struct {
	Replace string `pulumi:"replace"`
	Search  string `pulumi:"search"`
	Text    string `pulumi:"text"`
}

type ReplaceResult struct {
	Result string `pulumi:"result"`
}

func ReplaceOutput(ctx *pulumi.Context, args ReplaceOutputArgs, opts ...pulumi.InvokeOption) ReplaceResultOutput {
	outputResult := pulumix.ApplyErr[*ReplaceArgs](args.ToOutput(), func(plainArgs *ReplaceArgs) (*ReplaceResult, error) {
		return Replace(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[ReplaceResultOutput, *ReplaceResult](outputResult)
}

type ReplaceOutputArgs struct {
	Replace pulumix.Input[string] `pulumi:"replace"`
	Search  pulumix.Input[string] `pulumi:"search"`
	Text    pulumix.Input[string] `pulumi:"text"`
}

func (args ReplaceOutputArgs) ToOutput() pulumix.Output[*ReplaceArgs] {
	allArgs := pulumix.All(
		args.Replace.ToOutput(context.Background()).AsAny(),
		args.Search.ToOutput(context.Background()).AsAny(),
		args.Text.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *ReplaceArgs {
		return &ReplaceArgs{
			Replace: resolvedArgs[0].(string),
			Search:  resolvedArgs[1].(string),
			Text:    resolvedArgs[2].(string),
		}
	})
}

type ReplaceResultOutput struct{ *pulumi.OutputState }

func (ReplaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplaceResult)(nil)).Elem()
}

func (o ReplaceResultOutput) ToOutput(context.Context) pulumix.Output[*ReplaceResult] {
	return pulumix.Output[*ReplaceResult]{
		OutputState: o.OutputState,
	}
}

func (o ReplaceResultOutput) Result() pulumix.Output[string] {
	return pulumix.Apply[*ReplaceResult](o, func(v *ReplaceResult) string { return v.Result })
}