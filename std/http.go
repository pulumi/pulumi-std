package std

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"golang.org/x/net/http/httpproxy"
)

type Http struct{}

type HttpArgs struct {
	Url              string             `pulumi:"url"`
	Method           string             `pulumi:"method,optional"`
	RequestHeaders   *map[string]string `pulumi:"requestHeaders,optional"`
	RequestBody      *string            `pulumi:"requestBody,optional"`
	RequestTimeoutMs *int               `pulumi:"requestTimeoutMs,optional"`
	Insecure         *bool              `pulumi:"insecure,optional"`
	CaCertPem        *string            `pulumi:"caCertPem,optional"`
	Attempts         *int               `pulumi:"attempts,optional"`
	MinDelayMs       *int               `pulumi:"minDelayMs,optional"`
	MaxDelayMs       *int               `pulumi:"maxDelayMs,optional"`
}

type HttpResult struct {
	Id                 string            `pulumi:"id"`
	StatusCode         int               `pulumi:"statusCode"`
	ResponseHeaders    map[string]string `pulumi:"responseHeaders"`
	ResponseBody       string            `pulumi:"responseBody"`
	ResponseBodyBase64 string            `pulumi:"responseBodyBase64"`
}

func (r *Http) Annotate(a infer.Annotator) {
	a.Describe(r, "Performs an HTTP request and returns the response.")
}

func (*Http) Call(ctx p.Context, args HttpArgs) (HttpResult, error) {
	tr, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return HttpResult{}, fmt.Errorf("Error http: Can't configure http transport")
	}

	// Prevent issues with multiple data source configurations modifying the shared transport.
	clonedTr := tr.Clone()

	// Prevent issues with tests caching the proxy configuration.
	clonedTr.Proxy = func(req *http.Request) (*url.URL, error) {
		return httpproxy.FromEnvironment().ProxyFunc()(req.URL)
	}

	// the default method is GET
	if args.Method == "" {
		args.Method = "GET"
	}

	if args.Insecure != nil {
		if clonedTr.TLSClientConfig == nil {
			clonedTr.TLSClientConfig = &tls.Config{}
		}
		clonedTr.TLSClientConfig.InsecureSkipVerify = *args.Insecure
	}

	// Use `ca_cert_pem` certificate pool
	if args.CaCertPem != nil {
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM([]byte(*args.CaCertPem)); !ok {
			errorMsg := "Error tls: Can't add the CA certificate to certificate pool."
			return HttpResult{}, fmt.Errorf(errorMsg)
		}

		if clonedTr.TLSClientConfig == nil {
			clonedTr.TLSClientConfig = &tls.Config{}
		}
		clonedTr.TLSClientConfig.RootCAs = caCertPool
	}

	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = clonedTr
	if args.RequestTimeoutMs != nil {
		timeout := *args.RequestTimeoutMs
		retryClient.HTTPClient.Timeout = time.Duration(timeout) * time.Millisecond
	}

	// Configure retry settings
	if args.Attempts != nil {
		retryClient.RetryMax = *args.Attempts
	}

	if args.MinDelayMs != nil {
		minDelayMs := *args.MinDelayMs
		retryClient.RetryWaitMin = time.Duration(minDelayMs) * time.Millisecond
	}

	if args.MaxDelayMs != nil {
		maxDelayMs := *args.MaxDelayMs
		retryClient.RetryWaitMax = time.Duration(maxDelayMs) * time.Millisecond
	}

	requestBody := ""
	if args.RequestBody != nil {
		requestBody = *args.RequestBody
	}
	request, err := retryablehttp.NewRequestWithContext(ctx,
		args.Method,
		args.Url,
		strings.NewReader(requestBody))

	if err != nil {
		return HttpResult{}, fmt.Errorf("Can't create http request: %w", err)
	}

	if args.RequestHeaders != nil {
		for headerName, headerValue := range *args.RequestHeaders {
			request.Header.Set(headerName, headerValue)
		}
	}

	response, err := retryClient.Do(request)
	if err != nil {
		return HttpResult{}, fmt.Errorf("Can't send http request: %w", err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		return HttpResult{}, fmt.Errorf("Can't read http response body: %w", err)
	}

	if !utf8.Valid(bytes) {
		return HttpResult{}, fmt.Errorf("Error http: Response body is not valid UTF-8")
	}

	responseBody := string(bytes)
	responseBodyBase64Std := base64.StdEncoding.EncodeToString(bytes)
	responseHeaders := make(map[string]string)
	for k, v := range response.Header {
		// Concatenate according to RFC9110 https://www.rfc-editor.org/rfc/rfc9110.html#section-5.2
		responseHeaders[k] = strings.Join(v, ", ")
	}

	return HttpResult{
		Id:                 args.Url,
		StatusCode:         response.StatusCode,
		ResponseHeaders:    responseHeaders,
		ResponseBody:       responseBody,
		ResponseBodyBase64: responseBodyBase64Std,
	}, nil
}
