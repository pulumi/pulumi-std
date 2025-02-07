// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.Objects;


public final class LengthPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final LengthPlainArgs Empty = new LengthPlainArgs();

    @Import(name="input", required=true)
    private Object input;

    public Object input() {
        return this.input;
    }

    private LengthPlainArgs() {}

    private LengthPlainArgs(LengthPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LengthPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LengthPlainArgs $;

        public Builder() {
            $ = new LengthPlainArgs();
        }

        public Builder(LengthPlainArgs defaults) {
            $ = new LengthPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Object input) {
            $.input = input;
            return this;
        }

        public LengthPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("LengthPlainArgs", "input");
            }
            return $;
        }
    }

}
