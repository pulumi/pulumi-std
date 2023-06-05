// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the total sum of the elements of the input list.
 */
export function sum(args: SumArgs, opts?: pulumi.InvokeOptions): Promise<SumResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:sum", {
        "input": args.input,
    }, opts);
}

export interface SumArgs {
    input: number[];
}

export interface SumResult {
    readonly result: number;
}
/**
 * Returns the total sum of the elements of the input list.
 */
export function sumOutput(args: SumOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SumResult> {
    return pulumi.output(args).apply((a: any) => sum(a, opts))
}

export interface SumOutputArgs {
    input: pulumi.Input<pulumi.Input<number>[]>;
}
