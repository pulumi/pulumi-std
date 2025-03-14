// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the set of elements that all the input sets have in common.
 */
export function setintersection(args: SetintersectionArgs, opts?: pulumi.InvokeOptions): Promise<SetintersectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:setintersection", {
        "input": args.input,
    }, opts);
}

export interface SetintersectionArgs {
    input: any[][];
}

export interface SetintersectionResult {
    readonly result: any[];
}
/**
 * Returns the set of elements that all the input sets have in common.
 */
export function setintersectionOutput(args: SetintersectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<SetintersectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:setintersection", {
        "input": args.input,
    }, opts);
}

export interface SetintersectionOutputArgs {
    input: pulumi.Input<pulumi.Input<any[]>[]>;
}
