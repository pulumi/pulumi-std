// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Generates a list of numbers using a start value, a limit value, and a step value.
 * Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
 * depending on whether limit is greater than or less than start.
 */
export function range(args: RangeArgs, opts?: pulumi.InvokeOptions): Promise<RangeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:range", {
        "limit": args.limit,
        "start": args.start,
        "step": args.step,
    }, opts);
}

export interface RangeArgs {
    limit: number;
    start?: number;
    step?: number;
}

export interface RangeResult {
    readonly result: number[];
}
/**
 * Generates a list of numbers using a start value, a limit value, and a step value.
 * Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
 * depending on whether limit is greater than or less than start.
 */
export function rangeOutput(args: RangeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<RangeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:range", {
        "limit": args.limit,
        "start": args.start,
        "step": args.step,
    }, opts);
}

export interface RangeOutputArgs {
    limit: pulumi.Input<number>;
    start?: pulumi.Input<number>;
    step?: pulumi.Input<number>;
}
