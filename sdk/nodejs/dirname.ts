// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns all but the last element of path, typically the path's directory.
 */
export function dirname(args: DirnameArgs, opts?: pulumi.InvokeOptions): Promise<DirnameResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:dirname", {
        "input": args.input,
    }, opts);
}

export interface DirnameArgs {
    input: string;
}

export interface DirnameResult {
    readonly result: string;
}
/**
 * Returns all but the last element of path, typically the path's directory.
 */
export function dirnameOutput(args: DirnameOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<DirnameResult> {
    return pulumi.output(args).apply((a: any) => dirname(a, opts))
}

export interface DirnameOutputArgs {
    input: pulumi.Input<string>;
}
