// This file was auto-generated by Fern from our API Definition.

package activities

import (
	context "context"
	fmt "fmt"
	fern "github.com/mixed-file-directory/fern"
	core "github.com/mixed-file-directory/fern/core"
	giftcards "github.com/mixed-file-directory/fern/giftcards"
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

// Lists gift card activities. By default, you get gift card activities for all
// gift cards in the seller's account. You can optionally specify query parameters to
// filter the list. For example, you can get a list of gift card activities for a gift card,
// for all gift cards in a specific region, or for activities within a time window.
func (c *Client) List(
	ctx context.Context,
	request *giftcards.ActivitiesListRequest,
	opts ...option.RequestOption,
) (*core.Page[*fern.GiftCardActivity], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/gift-cards/activities"

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
	readPageResponse := func(response *fern.ListGiftCardActivitiesResponse) *core.PageResponse[*string, *fern.GiftCardActivity] {
		next := response.Cursor
		results := response.GiftCardActivities
		return &core.PageResponse[*string, *fern.GiftCardActivity]{
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

// Creates a gift card activity to manage the balance or state of a [gift card](entity:GiftCard).
// For example, you create an `ACTIVATE` activity to activate a gift card with an initial balance
// before the gift card can be used.
func (c *Client) Create(
	ctx context.Context,
	request *giftcards.CreateGiftCardActivityRequest,
	opts ...option.RequestOption,
) (*fern.CreateGiftCardActivityResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/gift-cards/activities"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.CreateGiftCardActivityResponse
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
