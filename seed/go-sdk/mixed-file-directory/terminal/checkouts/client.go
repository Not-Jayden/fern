// This file was auto-generated by Fern from our API Definition.

package checkouts

import (
	context "context"
	fern "github.com/mixed-file-directory/fern"
	core "github.com/mixed-file-directory/fern/core"
	option "github.com/mixed-file-directory/fern/option"
	terminal "github.com/mixed-file-directory/fern/terminal"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Creates a Terminal checkout request and sends it to the specified device to take a payment
// for the requested amount.
func (c *Client) Create(
	ctx context.Context,
	request *terminal.CreateTerminalCheckoutRequest,
	opts ...option.RequestOption,
) (*fern.CreateTerminalCheckoutResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/terminals/checkouts"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.CreateTerminalCheckoutResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns a filtered list of Terminal checkout requests created by the application making the request. Only Terminal checkout requests created for the merchant scoped to the OAuth token are returned. Terminal checkout requests are available for 30 days.
func (c *Client) Search(
	ctx context.Context,
	request *terminal.SearchTerminalCheckoutsRequest,
	opts ...option.RequestOption,
) (*fern.SearchTerminalCheckoutsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/terminals/checkouts/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.SearchTerminalCheckoutsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieves a Terminal checkout request by `checkout_id`. Terminal checkout requests are available for 30 days.
func (c *Client) Get(
	ctx context.Context,
	// The unique ID for the desired `TerminalCheckout`.
	checkoutId string,
	opts ...option.RequestOption,
) (*fern.GetTerminalCheckoutResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/terminals/checkouts/%v", checkoutId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.GetTerminalCheckoutResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Cancels a Terminal checkout request if the status of the request permits it.
func (c *Client) Cancel(
	ctx context.Context,
	// The unique ID for the desired `TerminalCheckout`.
	checkoutId string,
	request *terminal.CancelTerminalCheckoutRequest,
	opts ...option.RequestOption,
) (*fern.CancelTerminalCheckoutResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/terminals/checkouts/%v/cancel", checkoutId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.CancelTerminalCheckoutResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
