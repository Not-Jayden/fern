// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/unknown/fern/core"
	internal "github.com/unknown/fern/internal"
	option "github.com/unknown/fern/option"
	unknown "github.com/unknown/fern/unknown"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Unknown *unknown.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:  options.ToHeader(),
		Unknown: unknown.NewClient(opts...),
	}
}
