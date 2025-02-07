// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BcryptPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final BcryptPlainArgs Empty = new BcryptPlainArgs();

    @Import(name="cost")
    private @Nullable Integer cost;

    public Optional<Integer> cost() {
        return Optional.ofNullable(this.cost);
    }

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private BcryptPlainArgs() {}

    private BcryptPlainArgs(BcryptPlainArgs $) {
        this.cost = $.cost;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BcryptPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BcryptPlainArgs $;

        public Builder() {
            $ = new BcryptPlainArgs();
        }

        public Builder(BcryptPlainArgs defaults) {
            $ = new BcryptPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder cost(@Nullable Integer cost) {
            $.cost = cost;
            return this;
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public BcryptPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("BcryptPlainArgs", "input");
            }
            return $;
        }
    }

}
