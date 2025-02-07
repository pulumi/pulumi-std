// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.util.Objects;


public final class LogPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final LogPlainArgs Empty = new LogPlainArgs();

    @Import(name="base", required=true)
    private Double base;

    public Double base() {
        return this.base;
    }

    @Import(name="input", required=true)
    private Double input;

    public Double input() {
        return this.input;
    }

    private LogPlainArgs() {}

    private LogPlainArgs(LogPlainArgs $) {
        this.base = $.base;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogPlainArgs $;

        public Builder() {
            $ = new LogPlainArgs();
        }

        public Builder(LogPlainArgs defaults) {
            $ = new LogPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder base(Double base) {
            $.base = base;
            return this;
        }

        public Builder input(Double input) {
            $.input = input;
            return this;
        }

        public LogPlainArgs build() {
            if ($.base == null) {
                throw new MissingRequiredPropertyException("LogPlainArgs", "base");
            }
            if ($.input == null) {
                throw new MissingRequiredPropertyException("LogPlainArgs", "input");
            }
            return $;
        }
    }

}
