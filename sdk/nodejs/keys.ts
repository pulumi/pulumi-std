// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a lexically sorted list of the map keys.
 */
export function keys(args: KeysArgs, opts?: pulumi.InvokeOptions): Promise<KeysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:keys", {
        "input": args.input,
    }, opts);
}

export interface KeysArgs {
    input: {[key: string]: any};
}

export interface KeysResult {
    readonly result: string[];
}
/**
 * Returns a lexically sorted list of the map keys.
 */
export function keysOutput(args: KeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<KeysResult> {
    return pulumi.output(args).apply((a: any) => keys(a, opts))
}

export interface KeysOutputArgs {
    input: pulumi.Input<{[key: string]: any}>;
}
