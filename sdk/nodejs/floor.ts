// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the greatest integer value less than or equal to the argument.
 */
export function floor(args: FloorArgs, opts?: pulumi.InvokeOptions): Promise<FloorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:floor", {
        "input": args.input,
    }, opts);
}

export interface FloorArgs {
    input: number;
}

export interface FloorResult {
    readonly result: number;
}
/**
 * Returns the greatest integer value less than or equal to the argument.
 */
export function floorOutput(args: FloorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FloorResult> {
    return pulumi.output(args).apply((a: any) => floor(a, opts))
}

export interface FloorOutputArgs {
    input: pulumi.Input<number>;
}
