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


public final class ChunklistArgs extends com.pulumi.resources.InvokeArgs {

    public static final ChunklistArgs Empty = new ChunklistArgs();

    @Import(name="input", required=true)
    private Output<List<Object>> input;

    public Output<List<Object>> input() {
        return this.input;
    }

    @Import(name="size", required=true)
    private Output<Integer> size;

    public Output<Integer> size() {
        return this.size;
    }

    private ChunklistArgs() {}

    private ChunklistArgs(ChunklistArgs $) {
        this.input = $.input;
        this.size = $.size;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChunklistArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChunklistArgs $;

        public Builder() {
            $ = new ChunklistArgs();
        }

        public Builder(ChunklistArgs defaults) {
            $ = new ChunklistArgs(Objects.requireNonNull(defaults));
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

        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        public ChunklistArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("ChunklistArgs", "input");
            }
            if ($.size == null) {
                throw new MissingRequiredPropertyException("ChunklistArgs", "size");
            }
            return $;
        }
    }

}
