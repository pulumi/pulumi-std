// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class StartswithPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final StartswithPlainArgs Empty = new StartswithPlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    @Import(name="prefix", required=true)
    private String prefix;

    public String prefix() {
        return this.prefix;
    }

    private StartswithPlainArgs() {}

    private StartswithPlainArgs(StartswithPlainArgs $) {
        this.input = $.input;
        this.prefix = $.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StartswithPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StartswithPlainArgs $;

        public Builder() {
            $ = new StartswithPlainArgs();
        }

        public Builder(StartswithPlainArgs defaults) {
            $ = new StartswithPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Builder prefix(String prefix) {
            $.prefix = prefix;
            return this;
        }

        public StartswithPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("StartswithPlainArgs", "input");
            }
            if ($.prefix == null) {
                throw new MissingRequiredPropertyException("StartswithPlainArgs", "prefix");
            }
            return $;
        }
    }

}
