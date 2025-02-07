// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.Objects;


public final class ToboolArgs extends com.pulumi.resources.InvokeArgs {

    public static final ToboolArgs Empty = new ToboolArgs();

    @Import(name="input", required=true)
    private Output<Object> input;

    public Output<Object> input() {
        return this.input;
    }

    private ToboolArgs() {}

    private ToboolArgs(ToboolArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ToboolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ToboolArgs $;

        public Builder() {
            $ = new ToboolArgs();
        }

        public Builder(ToboolArgs defaults) {
            $ = new ToboolArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Object> input) {
            $.input = input;
            return this;
        }

        public Builder input(Object input) {
            return input(Output.of(input));
        }

        public ToboolArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("ToboolArgs", "input");
            }
            return $;
        }
    }

}
