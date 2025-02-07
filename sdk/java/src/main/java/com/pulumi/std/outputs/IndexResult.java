// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class IndexResult {
    private Integer result;

    private IndexResult() {}
    public Integer result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IndexResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer result;
        public Builder() {}
        public Builder(IndexResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Integer result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("IndexResult", "result");
            }
            this.result = result;
            return this;
        }
        public IndexResult build() {
            final var _resultValue = new IndexResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
