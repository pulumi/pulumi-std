// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;


public final class TransposeArgs extends com.pulumi.resources.InvokeArgs {

    public static final TransposeArgs Empty = new TransposeArgs();

    @Import(name="input", required=true)
    private Output<Map<String,List<String>>> input;

    public Output<Map<String,List<String>>> input() {
        return this.input;
    }

    private TransposeArgs() {}

    private TransposeArgs(TransposeArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransposeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransposeArgs $;

        public Builder() {
            $ = new TransposeArgs();
        }

        public Builder(TransposeArgs defaults) {
            $ = new TransposeArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<Map<String,List<String>>> input) {
            $.input = input;
            return this;
        }

        public Builder input(Map<String,List<String>> input) {
            return input(Output.of(input));
        }

        public TransposeArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}