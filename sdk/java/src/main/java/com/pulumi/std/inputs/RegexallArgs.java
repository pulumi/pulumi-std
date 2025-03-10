// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class RegexallArgs extends com.pulumi.resources.InvokeArgs {

    public static final RegexallArgs Empty = new RegexallArgs();

    @Import(name="pattern", required=true)
    private Output<String> pattern;

    public Output<String> pattern() {
        return this.pattern;
    }

    @Import(name="string", required=true)
    private Output<String> string;

    public Output<String> string() {
        return this.string;
    }

    private RegexallArgs() {}

    private RegexallArgs(RegexallArgs $) {
        this.pattern = $.pattern;
        this.string = $.string;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegexallArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegexallArgs $;

        public Builder() {
            $ = new RegexallArgs();
        }

        public Builder(RegexallArgs defaults) {
            $ = new RegexallArgs(Objects.requireNonNull(defaults));
        }

        public Builder pattern(Output<String> pattern) {
            $.pattern = pattern;
            return this;
        }

        public Builder pattern(String pattern) {
            return pattern(Output.of(pattern));
        }

        public Builder string(Output<String> string) {
            $.string = string;
            return this;
        }

        public Builder string(String string) {
            return string(Output.of(string));
        }

        public RegexallArgs build() {
            if ($.pattern == null) {
                throw new MissingRequiredPropertyException("RegexallArgs", "pattern");
            }
            if ($.string == null) {
                throw new MissingRequiredPropertyException("RegexallArgs", "string");
            }
            return $;
        }
    }

}
