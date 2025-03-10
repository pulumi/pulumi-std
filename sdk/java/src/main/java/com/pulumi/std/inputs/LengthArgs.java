// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.Objects;


public final class LengthArgs extends com.pulumi.resources.InvokeArgs {

    public static final LengthArgs Empty = new LengthArgs();

    @Import(name="input", required=true)
    private Output<Object> input;

    public Output<Object> input() {
        return this.input;
    }

    private LengthArgs() {}

    private LengthArgs(LengthArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LengthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LengthArgs $;

        public Builder() {
            $ = new LengthArgs();
        }

        public Builder(LengthArgs defaults) {
            $ = new LengthArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Object> input) {
            $.input = input;
            return this;
        }

        public Builder input(Object input) {
            return input(Output.of(input));
        }

        public LengthArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("LengthArgs", "input");
            }
            return $;
        }
    }

}
