// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class TrimArgs extends com.pulumi.resources.InvokeArgs {

    public static final TrimArgs Empty = new TrimArgs();

    @Import(name="cutset", required=true)
    private Output<String> cutset;

    public Output<String> cutset() {
        return this.cutset;
    }

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private TrimArgs() {}

    private TrimArgs(TrimArgs $) {
        this.cutset = $.cutset;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrimArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrimArgs $;

        public Builder() {
            $ = new TrimArgs();
        }

        public Builder(TrimArgs defaults) {
            $ = new TrimArgs(Objects.requireNonNull(defaults));
        }

        public Builder cutset(Output<String> cutset) {
            $.cutset = cutset;
            return this;
        }

        public Builder cutset(String cutset) {
            return cutset(Output.of(cutset));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public TrimArgs build() {
            $.cutset = Objects.requireNonNull($.cutset, "expected parameter 'cutset' to be non-null");
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
