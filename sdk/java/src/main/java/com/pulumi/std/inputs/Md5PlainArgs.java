// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class Md5PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Md5PlainArgs Empty = new Md5PlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Md5PlainArgs() {}

    private Md5PlainArgs(Md5PlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Md5PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Md5PlainArgs $;

        public Builder() {
            $ = new Md5PlainArgs();
        }

        public Builder(Md5PlainArgs defaults) {
            $ = new Md5PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Md5PlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("Md5PlainArgs", "input");
            }
            return $;
        }
    }

}
