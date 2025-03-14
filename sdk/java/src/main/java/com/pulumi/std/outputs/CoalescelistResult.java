// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.util.List;
import java.util.Objects;

@CustomType
public final class CoalescelistResult {
    private List<Object> result;

    private CoalescelistResult() {}
    public List<Object> result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CoalescelistResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Object> result;
        public Builder() {}
        public Builder(CoalescelistResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(List<Object> result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("CoalescelistResult", "result");
            }
            this.result = result;
            return this;
        }
        public Builder result(Object... result) {
            return result(List.of(result));
        }
        public CoalescelistResult build() {
            final var _resultValue = new CoalescelistResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
