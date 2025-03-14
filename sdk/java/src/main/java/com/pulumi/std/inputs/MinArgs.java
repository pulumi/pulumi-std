// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.List;
import java.util.Objects;


public final class MinArgs extends com.pulumi.resources.InvokeArgs {

    public static final MinArgs Empty = new MinArgs();

    @Import(name="input", required=true)
    private Output<List<Double>> input;

    public Output<List<Double>> input() {
        return this.input;
    }

    private MinArgs() {}

    private MinArgs(MinArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MinArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MinArgs $;

        public Builder() {
            $ = new MinArgs();
        }

        public Builder(MinArgs defaults) {
            $ = new MinArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<List<Double>> input) {
            $.input = input;
            return this;
        }

        public Builder input(List<Double> input) {
            return input(Output.of(input));
        }

        public Builder input(Double... input) {
            return input(List.of(input));
        }

        public MinArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("MinArgs", "input");
            }
            return $;
        }
    }

}
