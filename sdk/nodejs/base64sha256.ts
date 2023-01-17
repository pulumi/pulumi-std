// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a base64-encoded representation of raw SHA-256 sum of the given string.
 * This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
 */
export function base64sha256(args: Base64sha256Args, opts?: pulumi.InvokeOptions): Promise<Base64sha256Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:base64sha256", {
        "input": args.input,
    }, opts);
}

export interface Base64sha256Args {
    input: string;
}

export interface Base64sha256Result {
    readonly result: string;
}
/**
 * Returns a base64-encoded representation of raw SHA-256 sum of the given string.
 * This is not equivalent of base64encode(sha256(string)) since sha256() returns hexadecimal representation.
 */
export function base64sha256Output(args: Base64sha256OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Base64sha256Result> {
    return pulumi.output(args).apply((a: any) => base64sha256(a, opts))
}

export interface Base64sha256OutputArgs {
    input: pulumi.Input<string>;
}
