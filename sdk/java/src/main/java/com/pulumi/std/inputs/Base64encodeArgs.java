// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class Base64encodeArgs extends com.pulumi.resources.InvokeArgs {

    public static final Base64encodeArgs Empty = new Base64encodeArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Base64encodeArgs() {}

    private Base64encodeArgs(Base64encodeArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Base64encodeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Base64encodeArgs $;

        public Builder() {
            $ = new Base64encodeArgs();
        }

        public Builder(Base64encodeArgs defaults) {
            $ = new Base64encodeArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Base64encodeArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("Base64encodeArgs", "input");
            }
            return $;
        }
    }

}
