// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LookupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final LookupPlainArgs Empty = new LookupPlainArgs();

    @Import(name="default")
    private @Nullable Object default_;

    public Optional<Object> default_() {
        return Optional.ofNullable(this.default_);
    }

    @Import(name="key", required=true)
    private String key;

    public String key() {
        return this.key;
    }

    @Import(name="map", required=true)
    private Map<String,Object> map;

    public Map<String,Object> map() {
        return this.map;
    }

    private LookupPlainArgs() {}

    private LookupPlainArgs(LookupPlainArgs $) {
        this.default_ = $.default_;
        this.key = $.key;
        this.map = $.map;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LookupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LookupPlainArgs $;

        public Builder() {
            $ = new LookupPlainArgs();
        }

        public Builder(LookupPlainArgs defaults) {
            $ = new LookupPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder default_(@Nullable Object default_) {
            $.default_ = default_;
            return this;
        }

        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder map(Map<String,Object> map) {
            $.map = map;
            return this;
        }

        public LookupPlainArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("LookupPlainArgs", "key");
            }
            if ($.map == null) {
                throw new MissingRequiredPropertyException("LookupPlainArgs", "map");
            }
            return $;
        }
    }

}
