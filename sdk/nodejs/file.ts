// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Reads the contents of a file into the string.
 */
export function file(args: FileArgs, opts?: pulumi.InvokeOptions): Promise<FileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:file", {
        "input": args.input,
    }, opts);
}

export interface FileArgs {
    input: string;
}

export interface FileResult {
    readonly result: string;
}
/**
 * Reads the contents of a file into the string.
 */
export function fileOutput(args: FileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FileResult> {
    return pulumi.output(args).apply((a: any) => file(a, opts))
}

export interface FileOutputArgs {
    input: pulumi.Input<string>;
}
