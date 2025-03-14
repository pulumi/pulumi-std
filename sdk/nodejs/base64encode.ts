// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a base64-encoded representation of the given string.
 */
export function base64encode(args: Base64encodeArgs, opts?: pulumi.InvokeOptions): Promise<Base64encodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:base64encode", {
        "input": args.input,
    }, opts);
}

export interface Base64encodeArgs {
    input: string;
}

export interface Base64encodeResult {
    readonly result: string;
}
/**
 * Returns a base64-encoded representation of the given string.
 */
export function base64encodeOutput(args: Base64encodeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<Base64encodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:base64encode", {
        "input": args.input,
    }, opts);
}

export interface Base64encodeOutputArgs {
    input: pulumi.Input<string>;
}
