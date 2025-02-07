// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class TrimPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final TrimPlainArgs Empty = new TrimPlainArgs();

    @Import(name="cutset", required=true)
    private String cutset;

    public String cutset() {
        return this.cutset;
    }

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private TrimPlainArgs() {}

    private TrimPlainArgs(TrimPlainArgs $) {
        this.cutset = $.cutset;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrimPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrimPlainArgs $;

        public Builder() {
            $ = new TrimPlainArgs();
        }

        public Builder(TrimPlainArgs defaults) {
            $ = new TrimPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder cutset(String cutset) {
            $.cutset = cutset;
            return this;
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public TrimPlainArgs build() {
            if ($.cutset == null) {
                throw new MissingRequiredPropertyException("TrimPlainArgs", "cutset");
            }
            if ($.input == null) {
                throw new MissingRequiredPropertyException("TrimPlainArgs", "input");
            }
            return $;
        }
    }

}
