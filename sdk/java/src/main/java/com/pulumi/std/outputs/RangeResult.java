// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.List;
import java.util.Objects;

@CustomType
public final class RangeResult {
    private List<Double> result;

    private RangeResult() {}
    public List<Double> result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RangeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Double> result;
        public Builder() {}
        public Builder(RangeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(List<Double> result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("RangeResult", "result");
            }
            this.result = result;
            return this;
        }
        public Builder result(Double... result) {
            return result(List.of(result));
        }
        public RangeResult build() {
            final var _resultValue = new RangeResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
