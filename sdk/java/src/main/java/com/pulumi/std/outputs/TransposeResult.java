// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class TransposeResult {
    private Map<String,List<String>> result;

    private TransposeResult() {}
    public Map<String,List<String>> result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TransposeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,List<String>> result;
        public Builder() {}
        public Builder(TransposeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Map<String,List<String>> result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public TransposeResult build() {
            final var o = new TransposeResult();
            o.result = result;
            return o;
        }
    }
}
