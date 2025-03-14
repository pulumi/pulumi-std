// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class CoalesceArgs extends com.pulumi.resources.InvokeArgs {

    public static final CoalesceArgs Empty = new CoalesceArgs();

    @Import(name="input", required=true)
    private Output<List<Object>> input;

    public Output<List<Object>> input() {
        return this.input;
    }

    private CoalesceArgs() {}

    private CoalesceArgs(CoalesceArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CoalesceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CoalesceArgs $;

        public Builder() {
            $ = new CoalesceArgs();
        }

        public Builder(CoalesceArgs defaults) {
            $ = new CoalesceArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<List<Object>> input) {
            $.input = input;
            return this;
        }

        public Builder input(List<Object> input) {
            return input(Output.of(input));
        }

        public Builder input(Object... input) {
            return input(List.of(input));
        }

        public CoalesceArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("CoalesceArgs", "input");
            }
            return $;
        }
    }

}
