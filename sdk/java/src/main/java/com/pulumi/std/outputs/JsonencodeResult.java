// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class JsonencodeResult {
    private String result;

    private JsonencodeResult() {}
    public String result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JsonencodeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String result;
        public Builder() {}
        public Builder(JsonencodeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(String result) {
            if (result == null) {
              throw new MissingRequiredPropertyException("JsonencodeResult", "result");
            }
            this.result = result;
            return this;
        }
        public JsonencodeResult build() {
            final var _resultValue = new JsonencodeResult();
            _resultValue.result = result;
            return _resultValue;
        }
    }
}
