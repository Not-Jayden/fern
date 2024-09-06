// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/mixed-file-directory/fern/core"
	option "github.com/mixed-file-directory/fern/option"
	wagesetting "github.com/mixed-file-directory/fern/team/wagesetting"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	WageSetting *wagesetting.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:      options.ToHeader(),
		WageSetting: wagesetting.NewClient(opts...),
	}
}
