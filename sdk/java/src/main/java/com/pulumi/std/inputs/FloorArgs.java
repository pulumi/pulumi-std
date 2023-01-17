// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.util.Objects;


public final class FloorArgs extends com.pulumi.resources.InvokeArgs {

    public static final FloorArgs Empty = new FloorArgs();

    @Import(name="input", required=true)
    private Output<Double> input;

    public Output<Double> input() {
        return this.input;
    }

    private FloorArgs() {}

    private FloorArgs(FloorArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FloorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FloorArgs $;

        public Builder() {
            $ = new FloorArgs();
        }

        public Builder(FloorArgs defaults) {
            $ = new FloorArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Double> input) {
            $.input = input;
            return this;
        }

        public Builder input(Double input) {
            return input(Output.of(input));
        }

        public FloorArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}