// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.util.Objects;

@CustomType
public final class PowResult {
    private Double result;

    private PowResult() {}
    public Double result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PowResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Double result;
        public Builder() {}
        public Builder(PowResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Double result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public PowResult build() {
            final var o = new PowResult();
            o.result = result;
            return o;
        }
    }
}
