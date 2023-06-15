// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HttpArgs extends com.pulumi.resources.InvokeArgs {

    public static final HttpArgs Empty = new HttpArgs();

    @Import(name="attempts")
    private @Nullable Output<Integer> attempts;

    public Optional<Output<Integer>> attempts() {
        return Optional.ofNullable(this.attempts);
    }

    @Import(name="caCertPem")
    private @Nullable Output<String> caCertPem;

    public Optional<Output<String>> caCertPem() {
        return Optional.ofNullable(this.caCertPem);
    }

    @Import(name="insecure")
    private @Nullable Output<Boolean> insecure;

    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    @Import(name="maxDelayMs")
    private @Nullable Output<Integer> maxDelayMs;

    public Optional<Output<Integer>> maxDelayMs() {
        return Optional.ofNullable(this.maxDelayMs);
    }

    @Import(name="method")
    private @Nullable Output<String> method;

    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    @Import(name="minDelayMs")
    private @Nullable Output<Integer> minDelayMs;

    public Optional<Output<Integer>> minDelayMs() {
        return Optional.ofNullable(this.minDelayMs);
    }

    @Import(name="requestBody")
    private @Nullable Output<String> requestBody;

    public Optional<Output<String>> requestBody() {
        return Optional.ofNullable(this.requestBody);
    }

    @Import(name="requestHeaders")
    private @Nullable Output<Map<String,String>> requestHeaders;

    public Optional<Output<Map<String,String>>> requestHeaders() {
        return Optional.ofNullable(this.requestHeaders);
    }

    @Import(name="requestTimeoutMs")
    private @Nullable Output<Integer> requestTimeoutMs;

    public Optional<Output<Integer>> requestTimeoutMs() {
        return Optional.ofNullable(this.requestTimeoutMs);
    }

    @Import(name="url", required=true)
    private Output<String> url;

    public Output<String> url() {
        return this.url;
    }

    private HttpArgs() {}

    private HttpArgs(HttpArgs $) {
        this.attempts = $.attempts;
        this.caCertPem = $.caCertPem;
        this.insecure = $.insecure;
        this.maxDelayMs = $.maxDelayMs;
        this.method = $.method;
        this.minDelayMs = $.minDelayMs;
        this.requestBody = $.requestBody;
        this.requestHeaders = $.requestHeaders;
        this.requestTimeoutMs = $.requestTimeoutMs;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HttpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HttpArgs $;

        public Builder() {
            $ = new HttpArgs();
        }

        public Builder(HttpArgs defaults) {
            $ = new HttpArgs(Objects.requireNonNull(defaults));
        }

        public Builder attempts(@Nullable Output<Integer> attempts) {
            $.attempts = attempts;
            return this;
        }

        public Builder attempts(Integer attempts) {
            return attempts(Output.of(attempts));
        }

        public Builder caCertPem(@Nullable Output<String> caCertPem) {
            $.caCertPem = caCertPem;
            return this;
        }

        public Builder caCertPem(String caCertPem) {
            return caCertPem(Output.of(caCertPem));
        }

        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
        }

        public Builder maxDelayMs(@Nullable Output<Integer> maxDelayMs) {
            $.maxDelayMs = maxDelayMs;
            return this;
        }

        public Builder maxDelayMs(Integer maxDelayMs) {
            return maxDelayMs(Output.of(maxDelayMs));
        }

        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        public Builder method(String method) {
            return method(Output.of(method));
        }

        public Builder minDelayMs(@Nullable Output<Integer> minDelayMs) {
            $.minDelayMs = minDelayMs;
            return this;
        }

        public Builder minDelayMs(Integer minDelayMs) {
            return minDelayMs(Output.of(minDelayMs));
        }

        public Builder requestBody(@Nullable Output<String> requestBody) {
            $.requestBody = requestBody;
            return this;
        }

        public Builder requestBody(String requestBody) {
            return requestBody(Output.of(requestBody));
        }

        public Builder requestHeaders(@Nullable Output<Map<String,String>> requestHeaders) {
            $.requestHeaders = requestHeaders;
            return this;
        }

        public Builder requestHeaders(Map<String,String> requestHeaders) {
            return requestHeaders(Output.of(requestHeaders));
        }

        public Builder requestTimeoutMs(@Nullable Output<Integer> requestTimeoutMs) {
            $.requestTimeoutMs = requestTimeoutMs;
            return this;
        }

        public Builder requestTimeoutMs(Integer requestTimeoutMs) {
            return requestTimeoutMs(Output.of(requestTimeoutMs));
        }

        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public HttpArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
