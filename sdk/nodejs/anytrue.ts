// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns true if any of the elements in a given collection are true or \"true\".
 * It also returns false if the collection is empty.
 */
export function anytrue(args: AnytrueArgs, opts?: pulumi.InvokeOptions): Promise<AnytrueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:anytrue", {
        "input": args.input,
    }, opts);
}

export interface AnytrueArgs {
    input: any[];
}

export interface AnytrueResult {
    readonly result: boolean;
}
/**
 * Returns true if any of the elements in a given collection are true or \"true\".
 * It also returns false if the collection is empty.
 */
export function anytrueOutput(args: AnytrueOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<AnytrueResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("std:index:anytrue", {
        "input": args.input,
    }, opts);
}

export interface AnytrueOutputArgs {
    input: pulumi.Input<any[]>;
}
