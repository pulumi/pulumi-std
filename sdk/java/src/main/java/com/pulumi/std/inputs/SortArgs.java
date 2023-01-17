// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class SortArgs extends com.pulumi.resources.InvokeArgs {

    public static final SortArgs Empty = new SortArgs();

    @Import(name="input", required=true)
    private Output<List<String>> input;

    public Output<List<String>> input() {
        return this.input;
    }

    private SortArgs() {}

    private SortArgs(SortArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SortArgs $;

        public Builder() {
            $ = new SortArgs();
        }

        public Builder(SortArgs defaults) {
            $ = new SortArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<List<String>> input) {
            $.input = input;
            return this;
        }

        public Builder input(List<String> input) {
            return input(Output.of(input));
        }

        public Builder input(String... input) {
            return input(List.of(input));
        }

        public SortArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}