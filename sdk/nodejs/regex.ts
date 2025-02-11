// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the first matche of a regular expression in a string (including named or indexed groups).
 */
export function regex(args: RegexArgs, opts?: pulumi.InvokeOptions): Promise<RegexResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:regex", {
        "pattern": args.pattern,
        "string": args.string,
    }, opts);
}

export interface RegexArgs {
    pattern: string;
    string: string;
}

export interface RegexResult {
    readonly result: any;
}
/**
 * Returns the first matche of a regular expression in a string (including named or indexed groups).
 */
export function regexOutput(args: RegexOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<RegexResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:regex", {
        "pattern": args.pattern,
        "string": args.string,
    }, opts);
}

export interface RegexOutputArgs {
    pattern: pulumi.Input<string>;
    string: pulumi.Input<string>;
}
