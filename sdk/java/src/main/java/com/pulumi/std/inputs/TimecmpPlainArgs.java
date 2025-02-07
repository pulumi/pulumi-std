// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class TimecmpPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final TimecmpPlainArgs Empty = new TimecmpPlainArgs();

    @Import(name="timestampa", required=true)
    private String timestampa;

    public String timestampa() {
        return this.timestampa;
    }

    @Import(name="timestampb", required=true)
    private String timestampb;

    public String timestampb() {
        return this.timestampb;
    }

    private TimecmpPlainArgs() {}

    private TimecmpPlainArgs(TimecmpPlainArgs $) {
        this.timestampa = $.timestampa;
        this.timestampb = $.timestampb;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TimecmpPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TimecmpPlainArgs $;

        public Builder() {
            $ = new TimecmpPlainArgs();
        }

        public Builder(TimecmpPlainArgs defaults) {
            $ = new TimecmpPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder timestampa(String timestampa) {
            $.timestampa = timestampa;
            return this;
        }

        public Builder timestampb(String timestampb) {
            $.timestampb = timestampb;
            return this;
        }

        public TimecmpPlainArgs build() {
            if ($.timestampa == null) {
                throw new MissingRequiredPropertyException("TimecmpPlainArgs", "timestampa");
            }
            if ($.timestampb == null) {
                throw new MissingRequiredPropertyException("TimecmpPlainArgs", "timestampb");
            }
            return $;
        }
    }

}
