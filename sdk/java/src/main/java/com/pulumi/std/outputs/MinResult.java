// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.Objects;

@CustomType
public final class MinResult {
    private Double result;

    private MinResult() {}
    public Double result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MinResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Double result;
        public Builder() {}
        public Builder(MinResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Double result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("MinResult", "result");
            }
            this.result = result;
            return this;
        }
        public MinResult build() {
            final var _resultValue = new MinResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
