// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Converts all cased letters in the given string to uppercase.
 */
export function upper(args: UpperArgs, opts?: pulumi.InvokeOptions): Promise<UpperResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:upper", {
        "input": args.input,
    }, opts);
}

export interface UpperArgs {
    input: string;
}

export interface UpperResult {
    readonly result: string;
}
/**
 * Converts all cased letters in the given string to uppercase.
 */
export function upperOutput(args: UpperOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<UpperResult> {
    return pulumi.output(args).apply((a: any) => upper(a, opts))
}

export interface UpperOutputArgs {
    input: pulumi.Input<string>;
}
