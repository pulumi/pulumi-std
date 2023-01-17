// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Parses the given string as a representation of an integer in the specified base
 * and returns the resulting number. The base must be between 2 and 62 inclusive.
 * 	.
 */
export function parseint(args: ParseintArgs, opts?: pulumi.InvokeOptions): Promise<ParseintResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:parseint", {
        "base": args.base,
        "input": args.input,
    }, opts);
}

export interface ParseintArgs {
    base?: number;
    input: string;
}

export interface ParseintResult {
    readonly result: number;
}
/**
 * Parses the given string as a representation of an integer in the specified base
 * and returns the resulting number. The base must be between 2 and 62 inclusive.
 * 	.
 */
export function parseintOutput(args: ParseintOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ParseintResult> {
    return pulumi.output(args).apply((a: any) => parseint(a, opts))
}

export interface ParseintOutputArgs {
    base?: pulumi.Input<number>;
    input: pulumi.Input<string>;
}
