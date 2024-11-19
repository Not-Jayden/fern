// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/auth-environment-variables/fern/core"
	internal "github.com/auth-environment-variables/fern/internal"
	option "github.com/auth-environment-variables/fern/option"
	service "github.com/auth-environment-variables/fern/service"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Service *service.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.ApiKey == "" {
		options.ApiKey = os.Getenv("FERN_API_KEY")
	}
	if options.XAnotherHeader == "" {
		options.XAnotherHeader = os.Getenv("ANOTHER_ENV_VAR")
	}
	if options.XApiVersion == "" {
		options.XApiVersion = os.Getenv("VERSION")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:  options.ToHeader(),
		Service: service.NewClient(opts...),
	}
}
