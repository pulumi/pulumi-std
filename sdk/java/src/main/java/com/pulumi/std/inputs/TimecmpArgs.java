// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class TimecmpArgs extends com.pulumi.resources.InvokeArgs {

    public static final TimecmpArgs Empty = new TimecmpArgs();

    @Import(name="timestampa", required=true)
    private Output<String> timestampa;

    public Output<String> timestampa() {
        return this.timestampa;
    }

    @Import(name="timestampb", required=true)
    private Output<String> timestampb;

    public Output<String> timestampb() {
        return this.timestampb;
    }

    private TimecmpArgs() {}

    private TimecmpArgs(TimecmpArgs $) {
        this.timestampa = $.timestampa;
        this.timestampb = $.timestampb;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TimecmpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TimecmpArgs $;

        public Builder() {
            $ = new TimecmpArgs();
        }

        public Builder(TimecmpArgs defaults) {
            $ = new TimecmpArgs(Objects.requireNonNull(defaults));
        }

        public Builder timestampa(Output<String> timestampa) {
            $.timestampa = timestampa;
            return this;
        }

        public Builder timestampa(String timestampa) {
            return timestampa(Output.of(timestampa));
        }

        public Builder timestampb(Output<String> timestampb) {
            $.timestampb = timestampb;
            return this;
        }

        public Builder timestampb(String timestampb) {
            return timestampb(Output.of(timestampb));
        }

        public TimecmpArgs build() {
            if ($.timestampa == null) {
                throw new MissingRequiredPropertyException("TimecmpArgs", "timestampa");
            }
            if ($.timestampb == null) {
                throw new MissingRequiredPropertyException("TimecmpArgs", "timestampb");
            }
            return $;
        }
    }

}
