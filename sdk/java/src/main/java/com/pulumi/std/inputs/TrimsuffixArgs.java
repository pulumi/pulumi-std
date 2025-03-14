// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class TrimsuffixArgs extends com.pulumi.resources.InvokeArgs {

    public static final TrimsuffixArgs Empty = new TrimsuffixArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    @Import(name="suffix", required=true)
    private Output<String> suffix;

    public Output<String> suffix() {
        return this.suffix;
    }

    private TrimsuffixArgs() {}

    private TrimsuffixArgs(TrimsuffixArgs $) {
        this.input = $.input;
        this.suffix = $.suffix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrimsuffixArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrimsuffixArgs $;

        public Builder() {
            $ = new TrimsuffixArgs();
        }

        public Builder(TrimsuffixArgs defaults) {
            $ = new TrimsuffixArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Builder suffix(Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        public TrimsuffixArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("TrimsuffixArgs", "input");
            }
            if ($.suffix == null) {
                throw new MissingRequiredPropertyException("TrimsuffixArgs", "suffix");
            }
            return $;
        }
    }

}
