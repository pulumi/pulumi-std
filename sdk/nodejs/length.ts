// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Determines the length of a given list, map, or string.
 */
export function length(args: LengthArgs, opts?: pulumi.InvokeOptions): Promise<LengthResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:length", {
        "input": args.input,
    }, opts);
}

export interface LengthArgs {
    input: any;
}

export interface LengthResult {
    readonly result: number;
}
/**
 * Determines the length of a given list, map, or string.
 */
export function lengthOutput(args: LengthOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<LengthResult> {
    return pulumi.output(args).apply((a: any) => length(a, opts))
}

export interface LengthOutputArgs {
    input: any;
}
