// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AbspathResult {
    private String result;

    private AbspathResult() {}
    public String result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AbspathResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String result;
        public Builder() {}
        public Builder(AbspathResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(String result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public AbspathResult build() {
            final var o = new AbspathResult();
            o.result = result;
            return o;
        }
    }
}
