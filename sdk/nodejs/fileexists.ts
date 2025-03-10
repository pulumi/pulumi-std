// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Determines whether a file exists at a given path.
 */
export function fileexists(args: FileexistsArgs, opts?: pulumi.InvokeOptions): Promise<FileexistsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:fileexists", {
        "input": args.input,
    }, opts);
}

export interface FileexistsArgs {
    input: string;
}

export interface FileexistsResult {
    readonly result: boolean;
}
/**
 * Determines whether a file exists at a given path.
 */
export function fileexistsOutput(args: FileexistsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<FileexistsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:fileexists", {
        "input": args.input,
    }, opts);
}

export interface FileexistsOutputArgs {
    input: pulumi.Input<string>;
}
