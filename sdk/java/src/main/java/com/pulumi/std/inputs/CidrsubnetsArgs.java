// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class CidrsubnetsArgs extends com.pulumi.resources.InvokeArgs {

    public static final CidrsubnetsArgs Empty = new CidrsubnetsArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    @Import(name="newbits", required=true)
    private Output<List<Integer>> newbits;

    public Output<List<Integer>> newbits() {
        return this.newbits;
    }

    private CidrsubnetsArgs() {}

    private CidrsubnetsArgs(CidrsubnetsArgs $) {
        this.input = $.input;
        this.newbits = $.newbits;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CidrsubnetsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CidrsubnetsArgs $;

        public Builder() {
            $ = new CidrsubnetsArgs();
        }

        public Builder(CidrsubnetsArgs defaults) {
            $ = new CidrsubnetsArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Builder newbits(Output<List<Integer>> newbits) {
            $.newbits = newbits;
            return this;
        }

        public Builder newbits(List<Integer> newbits) {
            return newbits(Output.of(newbits));
        }

        public Builder newbits(Integer... newbits) {
            return newbits(List.of(newbits));
        }

        public CidrsubnetsArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("CidrsubnetsArgs", "input");
            }
            if ($.newbits == null) {
                throw new MissingRequiredPropertyException("CidrsubnetsArgs", "newbits");
            }
            return $;
        }
    }

}
