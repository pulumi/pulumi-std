// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the smallest of the floats.
 */
export function min(args: MinArgs, opts?: pulumi.InvokeOptions): Promise<MinResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:min", {
        "input": args.input,
    }, opts);
}

export interface MinArgs {
    input: number[];
}

export interface MinResult {
    readonly result: number;
}
/**
 * Returns the smallest of the floats.
 */
export function minOutput(args: MinOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<MinResult> {
    return pulumi.output(args).apply((a: any) => min(a, opts))
}

export interface MinOutputArgs {
    input: pulumi.Input<pulumi.Input<number>[]>;
}
