// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Base64gzipArgs extends com.pulumi.resources.InvokeArgs {

    public static final Base64gzipArgs Empty = new Base64gzipArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Base64gzipArgs() {}

    private Base64gzipArgs(Base64gzipArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Base64gzipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Base64gzipArgs $;

        public Builder() {
            $ = new Base64gzipArgs();
        }

        public Builder(Base64gzipArgs defaults) {
            $ = new Base64gzipArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Base64gzipArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}