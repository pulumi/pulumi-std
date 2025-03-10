// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the union of 2 or more maps. The maps are consumed in the order provided,
 * and duplicate keys overwrite previous entries.
 */
export function merge(args: MergeArgs, opts?: pulumi.InvokeOptions): Promise<MergeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:merge", {
        "input": args.input,
    }, opts);
}

export interface MergeArgs {
    input: {[key: string]: any}[];
}

export interface MergeResult {
    readonly result: {[key: string]: any};
}
/**
 * Returns the union of 2 or more maps. The maps are consumed in the order provided,
 * and duplicate keys overwrite previous entries.
 */
export function mergeOutput(args: MergeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<MergeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:merge", {
        "input": args.input,
    }, opts);
}

export interface MergeOutputArgs {
    input: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
}
