// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class SubstrArgs extends com.pulumi.resources.InvokeArgs {

    public static final SubstrArgs Empty = new SubstrArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    @Import(name="length", required=true)
    private Output<Integer> length;

    public Output<Integer> length() {
        return this.length;
    }

    @Import(name="offset", required=true)
    private Output<Integer> offset;

    public Output<Integer> offset() {
        return this.offset;
    }

    private SubstrArgs() {}

    private SubstrArgs(SubstrArgs $) {
        this.input = $.input;
        this.length = $.length;
        this.offset = $.offset;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubstrArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubstrArgs $;

        public Builder() {
            $ = new SubstrArgs();
        }

        public Builder(SubstrArgs defaults) {
            $ = new SubstrArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Builder length(Output<Integer> length) {
            $.length = length;
            return this;
        }

        public Builder length(Integer length) {
            return length(Output.of(length));
        }

        public Builder offset(Output<Integer> offset) {
            $.offset = offset;
            return this;
        }

        public Builder offset(Integer offset) {
            return offset(Output.of(offset));
        }

        public SubstrArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("SubstrArgs", "input");
            }
            if ($.length == null) {
                throw new MissingRequiredPropertyException("SubstrArgs", "length");
            }
            if ($.offset == null) {
                throw new MissingRequiredPropertyException("SubstrArgs", "offset");
            }
            return $;
        }
    }

}
