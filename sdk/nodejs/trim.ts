// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Removes the specified set of characters from the start and end of the given string.
 */
export function trim(args: TrimArgs, opts?: pulumi.InvokeOptions): Promise<TrimResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:trim", {
        "cutset": args.cutset,
        "input": args.input,
    }, opts);
}

export interface TrimArgs {
    cutset: string;
    input: string;
}

export interface TrimResult {
    readonly result: string;
}
/**
 * Removes the specified set of characters from the start and end of the given string.
 */
export function trimOutput(args: TrimOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<TrimResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:trim", {
        "cutset": args.cutset,
        "input": args.input,
    }, opts);
}

export interface TrimOutputArgs {
    cutset: pulumi.Input<string>;
    input: pulumi.Input<string>;
}
