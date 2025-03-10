// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Removes empty and nil string elements from a list.
 */
export function compact(args: CompactArgs, opts?: pulumi.InvokeOptions): Promise<CompactResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:compact", {
        "input": args.input,
    }, opts);
}

export interface CompactArgs {
    input: any[];
}

export interface CompactResult {
    readonly result: string[];
}
/**
 * Removes empty and nil string elements from a list.
 */
export function compactOutput(args: CompactOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<CompactResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:compact", {
        "input": args.input,
    }, opts);
}

export interface CompactOutputArgs {
    input: pulumi.Input<any[]>;
}
