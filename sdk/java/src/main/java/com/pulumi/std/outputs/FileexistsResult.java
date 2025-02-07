// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class FileexistsResult {
    private Boolean result;

    private FileexistsResult() {}
    public Boolean result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FileexistsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean result;
        public Builder() {}
        public Builder(FileexistsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Boolean result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("FileexistsResult", "result");
            }
            this.result = result;
            return this;
        }
        public FileexistsResult build() {
            final var _resultValue = new FileexistsResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
