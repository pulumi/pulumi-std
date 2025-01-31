// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the first non-empty list from the given list of lists.
 */
export function coalescelist(args: CoalescelistArgs, opts?: pulumi.InvokeOptions): Promise<CoalescelistResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:coalescelist", {
        "input": args.input,
    }, opts);
}

export interface CoalescelistArgs {
    input: any[][];
}

export interface CoalescelistResult {
    readonly result: any[];
}
/**
 * Returns the first non-empty list from the given list of lists.
 */
export function coalescelistOutput(args: CoalescelistOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<CoalescelistResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:coalescelist", {
        "input": args.input,
    }, opts);
}

export interface CoalescelistOutputArgs {
    input: pulumi.Input<pulumi.Input<any[]>[]>;
}
