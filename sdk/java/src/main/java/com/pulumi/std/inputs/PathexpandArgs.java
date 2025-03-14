// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class PathexpandArgs extends com.pulumi.resources.InvokeArgs {

    public static final PathexpandArgs Empty = new PathexpandArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private PathexpandArgs() {}

    private PathexpandArgs(PathexpandArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PathexpandArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PathexpandArgs $;

        public Builder() {
            $ = new PathexpandArgs();
        }

        public Builder(PathexpandArgs defaults) {
            $ = new PathexpandArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public PathexpandArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("PathexpandArgs", "input");
            }
            return $;
        }
    }

}
