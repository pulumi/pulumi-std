// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Reads the contents of a file into a string and returns the base64-encoded SHA512 hash of it.
 */
export function filebase64sha512(args: Filebase64sha512Args, opts?: pulumi.InvokeOptions): Promise<Filebase64sha512Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:filebase64sha512", {
        "input": args.input,
    }, opts);
}

export interface Filebase64sha512Args {
    input: string;
}

export interface Filebase64sha512Result {
    readonly result: string;
}
/**
 * Reads the contents of a file into a string and returns the base64-encoded SHA512 hash of it.
 */
export function filebase64sha512Output(args: Filebase64sha512OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Filebase64sha512Result> {
    return pulumi.output(args).apply((a: any) => filebase64sha512(a, opts))
}

export interface Filebase64sha512OutputArgs {
    input: pulumi.Input<string>;
}
