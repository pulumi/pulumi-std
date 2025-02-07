// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class Base64gzipPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Base64gzipPlainArgs Empty = new Base64gzipPlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Base64gzipPlainArgs() {}

    private Base64gzipPlainArgs(Base64gzipPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Base64gzipPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Base64gzipPlainArgs $;

        public Builder() {
            $ = new Base64gzipPlainArgs();
        }

        public Builder(Base64gzipPlainArgs defaults) {
            $ = new Base64gzipPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Base64gzipPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("Base64gzipPlainArgs", "input");
            }
            return $;
        }
    }

}
