// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class RsadecryptPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final RsadecryptPlainArgs Empty = new RsadecryptPlainArgs();

    @Import(name="cipherText", required=true)
    private String cipherText;

    public String cipherText() {
        return this.cipherText;
    }

    @Import(name="key", required=true)
    private String key;

    public String key() {
        return this.key;
    }

    private RsadecryptPlainArgs() {}

    private RsadecryptPlainArgs(RsadecryptPlainArgs $) {
        this.cipherText = $.cipherText;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RsadecryptPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RsadecryptPlainArgs $;

        public Builder() {
            $ = new RsadecryptPlainArgs();
        }

        public Builder(RsadecryptPlainArgs defaults) {
            $ = new RsadecryptPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder cipherText(String cipherText) {
            $.cipherText = cipherText;
            return this;
        }

        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public RsadecryptPlainArgs build() {
            if ($.cipherText == null) {
                throw new MissingRequiredPropertyException("RsadecryptPlainArgs", "cipherText");
            }
            if ($.key == null) {
                throw new MissingRequiredPropertyException("RsadecryptPlainArgs", "key");
            }
            return $;
        }
    }

}
