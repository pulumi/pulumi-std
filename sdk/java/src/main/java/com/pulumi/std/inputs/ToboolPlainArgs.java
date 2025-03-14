// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.Objects;


public final class ToboolPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final ToboolPlainArgs Empty = new ToboolPlainArgs();

    @Import(name="input", required=true)
    private Object input;

    public Object input() {
        return this.input;
    }

    private ToboolPlainArgs() {}

    private ToboolPlainArgs(ToboolPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ToboolPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ToboolPlainArgs $;

        public Builder() {
            $ = new ToboolPlainArgs();
        }

        public Builder(ToboolPlainArgs defaults) {
            $ = new ToboolPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Object input) {
            $.input = input;
            return this;
        }

        public ToboolPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("ToboolPlainArgs", "input");
            }
            return $;
        }
    }

}
