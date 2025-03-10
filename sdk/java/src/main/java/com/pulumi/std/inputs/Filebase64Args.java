// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class Filebase64Args extends com.pulumi.resources.InvokeArgs {

    public static final Filebase64Args Empty = new Filebase64Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Filebase64Args() {}

    private Filebase64Args(Filebase64Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Filebase64Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Filebase64Args $;

        public Builder() {
            $ = new Filebase64Args();
        }

        public Builder(Filebase64Args defaults) {
            $ = new Filebase64Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Filebase64Args build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("Filebase64Args", "input");
            }
            return $;
        }
    }

}
