// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class ElementArgs extends com.pulumi.resources.InvokeArgs {

    public static final ElementArgs Empty = new ElementArgs();

    @Import(name="index", required=true)
    private Output<Integer> index;

    public Output<Integer> index() {
        return this.index;
    }

    @Import(name="input", required=true)
    private Output<List<Object>> input;

    public Output<List<Object>> input() {
        return this.input;
    }

    private ElementArgs() {}

    private ElementArgs(ElementArgs $) {
        this.index = $.index;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ElementArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ElementArgs $;

        public Builder() {
            $ = new ElementArgs();
        }

        public Builder(ElementArgs defaults) {
            $ = new ElementArgs(Objects.requireNonNull(defaults));
        }

        public Builder index(Output<Integer> index) {
            $.index = index;
            return this;
        }

        public Builder index(Integer index) {
            return index(Output.of(index));
        }

        public Builder input(Output<List<Object>> input) {
            $.input = input;
            return this;
        }

        public Builder input(List<Object> input) {
            return input(Output.of(input));
        }

        public Builder input(Object... input) {
            return input(List.of(input));
        }

        public ElementArgs build() {
            if ($.index == null) {
                throw new MissingRequiredPropertyException("ElementArgs", "index");
            }
            if ($.input == null) {
                throw new MissingRequiredPropertyException("ElementArgs", "input");
            }
            return $;
        }
    }

}
