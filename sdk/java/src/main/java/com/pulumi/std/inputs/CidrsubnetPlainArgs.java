// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class CidrsubnetPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final CidrsubnetPlainArgs Empty = new CidrsubnetPlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    @Import(name="netnum", required=true)
    private Integer netnum;

    public Integer netnum() {
        return this.netnum;
    }

    @Import(name="newbits", required=true)
    private Integer newbits;

    public Integer newbits() {
        return this.newbits;
    }

    private CidrsubnetPlainArgs() {}

    private CidrsubnetPlainArgs(CidrsubnetPlainArgs $) {
        this.input = $.input;
        this.netnum = $.netnum;
        this.newbits = $.newbits;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CidrsubnetPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CidrsubnetPlainArgs $;

        public Builder() {
            $ = new CidrsubnetPlainArgs();
        }

        public Builder(CidrsubnetPlainArgs defaults) {
            $ = new CidrsubnetPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Builder netnum(Integer netnum) {
            $.netnum = netnum;
            return this;
        }

        public Builder newbits(Integer newbits) {
            $.newbits = newbits;
            return this;
        }

        public CidrsubnetPlainArgs build() {
            if ($.input == null) {
                throw new MissingRequiredPropertyException("CidrsubnetPlainArgs", "input");
            }
            if ($.netnum == null) {
                throw new MissingRequiredPropertyException("CidrsubnetPlainArgs", "netnum");
            }
            if ($.newbits == null) {
                throw new MissingRequiredPropertyException("CidrsubnetPlainArgs", "newbits");
            }
            return $;
        }
    }

}
