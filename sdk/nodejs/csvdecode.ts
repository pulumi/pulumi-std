// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
 * 	The first line of the CSV data is interpreted as a "header" row: the values given
 * 	are used as the keys in the resulting maps.
 * 	Each subsequent line becomes a single map in the resulting list,
 * 	matching the keys from the header row with the given values by index.
 * 	All lines in the file must contain the same number of fields,
 * 	or this function will produce an error.
 * 	Follows the format defined in RFC 4180.
 */
export function csvdecode(args: CsvdecodeArgs, opts?: pulumi.InvokeOptions): Promise<CsvdecodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:csvdecode", {
        "input": args.input,
    }, opts);
}

export interface CsvdecodeArgs {
    input: string;
}

export interface CsvdecodeResult {
    readonly result: {[key: string]: string}[];
}
/**
 * Decodes a string containing CSV-formatted data and produces a list of maps representing that data.
 * 	The first line of the CSV data is interpreted as a "header" row: the values given
 * 	are used as the keys in the resulting maps.
 * 	Each subsequent line becomes a single map in the resulting list,
 * 	matching the keys from the header row with the given values by index.
 * 	All lines in the file must contain the same number of fields,
 * 	or this function will produce an error.
 * 	Follows the format defined in RFC 4180.
 */
export function csvdecodeOutput(args: CsvdecodeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<CsvdecodeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:csvdecode", {
        "input": args.input,
    }, opts);
}

export interface CsvdecodeOutputArgs {
    input: pulumi.Input<string>;
}
