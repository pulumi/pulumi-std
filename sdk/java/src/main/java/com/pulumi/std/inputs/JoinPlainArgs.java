// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class JoinPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final JoinPlainArgs Empty = new JoinPlainArgs();

    @Import(name="input", required=true)
    private List<String> input;

    public List<String> input() {
        return this.input;
    }

    @Import(name="separator", required=true)
    private String separator;

    public String separator() {
        return this.separator;
    }

    private JoinPlainArgs() {}

    private JoinPlainArgs(JoinPlainArgs $) {
        this.input = $.input;
        this.separator = $.separator;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JoinPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JoinPlainArgs $;

        public Builder() {
            $ = new JoinPlainArgs();
        }

        public Builder(JoinPlainArgs defaults) {
            $ = new JoinPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(List<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String... input) {
            return input(List.of(input));
        }

        public Builder separator(String separator) {
            $.separator = separator;
            return this;
        }

        public JoinPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("JoinPlainArgs", "input");
            }
            if ($.separator == null) {
                throw new MissingRequiredPropertyException("JoinPlainArgs", "separator");
            }
            return $;
        }
    }

}
