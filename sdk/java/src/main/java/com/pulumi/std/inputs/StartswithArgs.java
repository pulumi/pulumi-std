// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class StartswithArgs extends com.pulumi.resources.InvokeArgs {

    public static final StartswithArgs Empty = new StartswithArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    @Import(name="prefix", required=true)
    private Output<String> prefix;

    public Output<String> prefix() {
        return this.prefix;
    }

    private StartswithArgs() {}

    private StartswithArgs(StartswithArgs $) {
        this.input = $.input;
        this.prefix = $.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StartswithArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StartswithArgs $;

        public Builder() {
            $ = new StartswithArgs();
        }

        public Builder(StartswithArgs defaults) {
            $ = new StartswithArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Builder prefix(Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        public StartswithArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("StartswithArgs", "input");
            }
            if ($.prefix == null) {
                throw new MissingRequiredPropertyException("StartswithArgs", "prefix");
            }
            return $;
        }
    }

}
