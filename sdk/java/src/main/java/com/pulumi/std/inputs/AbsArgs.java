// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.Objects;


public final class AbsArgs extends com.pulumi.resources.InvokeArgs {

    public static final AbsArgs Empty = new AbsArgs();

    @Import(name="input", required=true)
    private Output<Double> input;

    public Output<Double> input() {
        return this.input;
    }

    private AbsArgs() {}

    private AbsArgs(AbsArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AbsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AbsArgs $;

        public Builder() {
            $ = new AbsArgs();
        }

        public Builder(AbsArgs defaults) {
            $ = new AbsArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Double> input) {
            $.input = input;
            return this;
        }

        public Builder input(Double input) {
            return input(Output.of(input));
        }

        public AbsArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("AbsArgs", "input");
            }
            return $;
        }
    }

}
