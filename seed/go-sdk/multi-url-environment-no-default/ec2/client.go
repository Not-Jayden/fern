// This file was auto-generated by Fern from our API Definition.

package ec2

import (
	context "context"
	fern "github.com/multi-url-environment-no-default/fern"
	core "github.com/multi-url-environment-no-default/fern/core"
	internal "github.com/multi-url-environment-no-default/fern/internal"
	option "github.com/multi-url-environment-no-default/fern/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

func (c *Client) BootInstance(
	ctx context.Context,
	request *fern.BootInstanceRequest,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://ec2.aws.com",
	)
	endpointURL := baseURL + "/ec2/boot"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
		},
	); err != nil {
		return err
	}
	return nil
}
