// This file was auto-generated by Fern from our API Definition.

package breaktypes

import (
	context "context"
	fern "github.com/mixed-file-directory/fern"
	core "github.com/mixed-file-directory/fern/core"
	labor "github.com/mixed-file-directory/fern/labor"
	option "github.com/mixed-file-directory/fern/option"
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

// Returns a single `BreakType` specified by `id`.
func (c *Client) Get(
	ctx context.Context,
	// The UUID for the `BreakType` being retrieved.
	id string,
	opts ...option.RequestOption,
) (*fern.GetBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/labor/break-types/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.GetBreakTypeResponse
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

// Updates an existing `BreakType`.
func (c *Client) Update(
	ctx context.Context,
	// The UUID for the `BreakType` being updated.
	id string,
	request *labor.UpdateBreakTypeRequest,
	opts ...option.RequestOption,
) (*fern.UpdateBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/labor/break-types/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.UpdateBreakTypeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPut,
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

// Deletes an existing `BreakType`.
//
// A `BreakType` can be deleted even if it is referenced from a `Shift`.
func (c *Client) Delete(
	ctx context.Context,
	// The UUID for the `BreakType` being deleted.
	id string,
	opts ...option.RequestOption,
) (*fern.DeleteBreakTypeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/labor/break-types/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.DeleteBreakTypeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodDelete,
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
