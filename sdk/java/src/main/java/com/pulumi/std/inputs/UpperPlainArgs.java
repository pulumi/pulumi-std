// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class UpperPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final UpperPlainArgs Empty = new UpperPlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private UpperPlainArgs() {}

    private UpperPlainArgs(UpperPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UpperPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UpperPlainArgs $;

        public Builder() {
            $ = new UpperPlainArgs();
        }

        public Builder(UpperPlainArgs defaults) {
            $ = new UpperPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public UpperPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("UpperPlainArgs", "input");
            }
            return $;
        }
    }

}
