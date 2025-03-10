// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class AbspathArgs extends com.pulumi.resources.InvokeArgs {

    public static final AbspathArgs Empty = new AbspathArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private AbspathArgs() {}

    private AbspathArgs(AbspathArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AbspathArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AbspathArgs $;

        public Builder() {
            $ = new AbspathArgs();
        }

        public Builder(AbspathArgs defaults) {
            $ = new AbspathArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public AbspathArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("AbspathArgs", "input");
            }
            return $;
        }
    }

}
