// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a sequence with the same elements but in reverse order.
 */
export function reverse(args: ReverseArgs, opts?: pulumi.InvokeOptions): Promise<ReverseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:reverse", {
        "input": args.input,
    }, opts);
}

export interface ReverseArgs {
    input: any[];
}

export interface ReverseResult {
    readonly result: any[];
}
/**
 * Returns a sequence with the same elements but in reverse order.
 */
export function reverseOutput(args: ReverseOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ReverseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:reverse", {
        "input": args.input,
    }, opts);
}

export interface ReverseOutputArgs {
    input: pulumi.Input<any[]>;
}
