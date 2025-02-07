// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.Objects;

@CustomType
public final class SignumResult {
    private Double result;

    private SignumResult() {}
    public Double result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SignumResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Double result;
        public Builder() {}
        public Builder(SignumResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Double result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("SignumResult", "result");
            }
            this.result = result;
            return this;
        }
        public SignumResult build() {
            final var _resultValue = new SignumResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
