// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ChompArgs extends com.pulumi.resources.InvokeArgs {

    public static final ChompArgs Empty = new ChompArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private ChompArgs() {}

    private ChompArgs(ChompArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChompArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChompArgs $;

        public Builder() {
            $ = new ChompArgs();
        }

        public Builder(ChompArgs defaults) {
            $ = new ChompArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public ChompArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("ChompArgs", "input");
            }
            return $;
        }
    }

}
