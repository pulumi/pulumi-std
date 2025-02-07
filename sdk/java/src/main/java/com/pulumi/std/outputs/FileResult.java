// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FileResult {
    private String result;

    private FileResult() {}
    public String result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FileResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String result;
        public Builder() {}
        public Builder(FileResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(String result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("FileResult", "result");
            }
            this.result = result;
            return this;
        }
        public FileResult build() {
            final var _resultValue = new FileResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
