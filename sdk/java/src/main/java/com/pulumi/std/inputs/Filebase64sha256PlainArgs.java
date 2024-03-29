// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Filebase64sha256PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Filebase64sha256PlainArgs Empty = new Filebase64sha256PlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Filebase64sha256PlainArgs() {}

    private Filebase64sha256PlainArgs(Filebase64sha256PlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Filebase64sha256PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Filebase64sha256PlainArgs $;

        public Builder() {
            $ = new Filebase64sha256PlainArgs();
        }

        public Builder(Filebase64sha256PlainArgs defaults) {
            $ = new Filebase64sha256PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Filebase64sha256PlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
