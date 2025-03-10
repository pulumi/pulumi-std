// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a hexadecimal representation of the SHA-512 hash of the given string.
 */
export function sha512(args: Sha512Args, opts?: pulumi.InvokeOptions): Promise<Sha512Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:sha512", {
        "input": args.input,
    }, opts);
}

export interface Sha512Args {
    input: string;
}

export interface Sha512Result {
    readonly result: string;
}
/**
 * Returns a hexadecimal representation of the SHA-512 hash of the given string.
 */
export function sha512Output(args: Sha512OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<Sha512Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:sha512", {
        "input": args.input,
    }, opts);
}

export interface Sha512OutputArgs {
    input: pulumi.Input<string>;
}
