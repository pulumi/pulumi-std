// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the given string with all of its Unicode characters in reverse order.
 */
export function strrev(args: StrrevArgs, opts?: pulumi.InvokeOptions): Promise<StrrevResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:strrev", {
        "input": args.input,
    }, opts);
}

export interface StrrevArgs {
    input: string;
}

export interface StrrevResult {
    readonly result: string;
}
/**
 * Returns the given string with all of its Unicode characters in reverse order.
 */
export function strrevOutput(args: StrrevOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<StrrevResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:strrev", {
        "input": args.input,
    }, opts);
}

export interface StrrevOutputArgs {
    input: pulumi.Input<string>;
}
