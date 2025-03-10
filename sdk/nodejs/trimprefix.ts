// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Removes the specified prefix from the start of the given string, if present.
 */
export function trimprefix(args: TrimprefixArgs, opts?: pulumi.InvokeOptions): Promise<TrimprefixResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:trimprefix", {
        "input": args.input,
        "prefix": args.prefix,
    }, opts);
}

export interface TrimprefixArgs {
    input: string;
    prefix: string;
}

export interface TrimprefixResult {
    readonly result: string;
}
/**
 * Removes the specified prefix from the start of the given string, if present.
 */
export function trimprefixOutput(args: TrimprefixOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<TrimprefixResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:trimprefix", {
        "input": args.input,
        "prefix": args.prefix,
    }, opts);
}

export interface TrimprefixOutputArgs {
    input: pulumi.Input<string>;
    prefix: pulumi.Input<string>;
}
