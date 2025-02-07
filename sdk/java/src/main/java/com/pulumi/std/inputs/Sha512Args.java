// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class Sha512Args extends com.pulumi.resources.InvokeArgs {

    public static final Sha512Args Empty = new Sha512Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Sha512Args() {}

    private Sha512Args(Sha512Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Sha512Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Sha512Args $;

        public Builder() {
            $ = new Sha512Args();
        }

        public Builder(Sha512Args defaults) {
            $ = new Sha512Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Sha512Args build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("Sha512Args", "input");
            }
            return $;
        }
    }

}
