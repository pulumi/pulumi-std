// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class LengthResult {
    private Integer result;

    private LengthResult() {}
    public Integer result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LengthResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer result;
        public Builder() {}
        public Builder(LengthResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Integer result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public LengthResult build() {
            final var o = new LengthResult();
            o.result = result;
            return o;
        }
    }
}
