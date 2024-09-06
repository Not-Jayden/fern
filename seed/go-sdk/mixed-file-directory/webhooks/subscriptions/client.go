// This file was auto-generated by Fern from our API Definition.

package subscriptions

import (
	context "context"
	fmt "fmt"
	fern "github.com/mixed-file-directory/fern"
	core "github.com/mixed-file-directory/fern/core"
	option "github.com/mixed-file-directory/fern/option"
	webhooks "github.com/mixed-file-directory/fern/webhooks"
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

// Lists all webhook subscriptions owned by your application.
func (c *Client) List(
	ctx context.Context,
	request *webhooks.SubscriptionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*fern.WebhookSubscription], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/webhooks/subscriptions"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*string]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &core.CallParams{
			URL:         nextURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    pageRequest.Response,
		}
	}
	readPageResponse := func(response *fern.ListWebhookSubscriptionsResponse) *core.PageResponse[*string, *fern.WebhookSubscription] {
		next := response.Cursor
		results := response.Subscriptions
		return &core.PageResponse[*string, *fern.WebhookSubscription]{
			Next:    next,
			Results: results,
		}
	}
	pager := core.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Creates a webhook subscription.
func (c *Client) Create(
	ctx context.Context,
	request *webhooks.CreateWebhookSubscriptionRequest,
	opts ...option.RequestOption,
) (*fern.CreateWebhookSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/webhooks/subscriptions"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.CreateWebhookSubscriptionResponse
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

// Retrieves a webhook subscription identified by its ID.
func (c *Client) Get(
	ctx context.Context,
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to retrieve.
	subscriptionId string,
	opts ...option.RequestOption,
) (*fern.GetWebhookSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/webhooks/subscriptions/%v", subscriptionId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.GetWebhookSubscriptionResponse
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

// Updates a webhook subscription.
func (c *Client) Update(
	ctx context.Context,
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
	subscriptionId string,
	request *webhooks.UpdateWebhookSubscriptionRequest,
	opts ...option.RequestOption,
) (*fern.UpdateWebhookSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/webhooks/subscriptions/%v", subscriptionId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.UpdateWebhookSubscriptionResponse
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

// Deletes a webhook subscription.
func (c *Client) Delete(
	ctx context.Context,
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to delete.
	subscriptionId string,
	opts ...option.RequestOption,
) (*fern.DeleteWebhookSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/webhooks/subscriptions/%v", subscriptionId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.DeleteWebhookSubscriptionResponse
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

// Updates a webhook subscription by replacing the existing signature key with a new one.
func (c *Client) UpdateSignatureKey(
	ctx context.Context,
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update.
	subscriptionId string,
	request *webhooks.UpdateWebhookSubscriptionSignatureKeyRequest,
	opts ...option.RequestOption,
) (*fern.UpdateWebhookSubscriptionSignatureKeyResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/webhooks/subscriptions/%v/signature-key", subscriptionId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.UpdateWebhookSubscriptionSignatureKeyResponse
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

// Tests a webhook subscription by sending a test event to the notification URL.
func (c *Client) Test(
	ctx context.Context,
	// [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to test.
	subscriptionId string,
	request *webhooks.TestWebhookSubscriptionRequest,
	opts ...option.RequestOption,
) (*fern.TestWebhookSubscriptionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/webhooks/subscriptions/%v/test", subscriptionId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.TestWebhookSubscriptionResponse
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
