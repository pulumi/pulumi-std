// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Converts its argument to a set value.
 */
export function toset(args: TosetArgs, opts?: pulumi.InvokeOptions): Promise<TosetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:toset", {
        "input": args.input,
    }, opts);
}

export interface TosetArgs {
    input: any[];
}

export interface TosetResult {
    readonly result: any[];
}
/**
 * Converts its argument to a set value.
 */
export function tosetOutput(args: TosetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TosetResult> {
    return pulumi.output(args).apply((a: any) => toset(a, opts))
}

export interface TosetOutputArgs {
    input: pulumi.Input<any[]>;
}