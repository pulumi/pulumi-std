// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class FileexistsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final FileexistsPlainArgs Empty = new FileexistsPlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private FileexistsPlainArgs() {}

    private FileexistsPlainArgs(FileexistsPlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileexistsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileexistsPlainArgs $;

        public Builder() {
            $ = new FileexistsPlainArgs();
        }

        public Builder(FileexistsPlainArgs defaults) {
            $ = new FileexistsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public FileexistsPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("FileexistsPlainArgs", "input");
            }
            return $;
        }
    }

}
