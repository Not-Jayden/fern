// This file was auto-generated by Fern from our API Definition.

package users

import (
	context "context"
	fmt "fmt"
	uuid "github.com/google/uuid"
	fern "github.com/pagination/fern"
	core "github.com/pagination/fern/core"
	option "github.com/pagination/fern/option"
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

func (c *Client) ListWithCursorPagination(
	ctx context.Context,
	request *fern.ListUsersCursorPaginationRequest,
	opts ...option.RequestOption,
) (*core.Page[*fern.User], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*string]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("starting_after", fmt.Sprintf("%v", *pageRequest.Cursor))
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
	readPageResponse := func(response *fern.ListUsersPaginationResponse) *core.PageResponse[*string, *fern.User] {
		var next string
		if response.Page != nil && response.Page.Next != nil {
			next = response.Page.Next.StartingAfter
		}
		results := response.Data
		return &core.PageResponse[*string, *fern.User]{
			Next:    &next,
			Results: results,
		}
	}
	pager := core.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.StartingAfter)
}

func (c *Client) ListWithOffsetPagination(
	ctx context.Context,
	request *fern.ListUsersOffsetPaginationRequest,
	opts ...option.RequestOption,
) (*core.Page[*fern.User], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*int]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("page", fmt.Sprintf("%v", *pageRequest.Cursor))
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
	next := 1
	if request.Page != nil {
		next = *request.Page
	}
	readPageResponse := func(response *fern.ListUsersPaginationResponse) *core.PageResponse[*int, *fern.User] {
		next += 1
		results := response.Data
		return &core.PageResponse[*int, *fern.User]{
			Next:    &next,
			Results: results,
		}
	}
	pager := core.NewOffsetPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, &next)
}

func (c *Client) ListWithExtendedResults(
	ctx context.Context,
	request *fern.ListUsersExtendedRequest,
	opts ...option.RequestOption,
) (*core.Page[*fern.User], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*uuid.UUID]) *core.CallParams {
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
	readPageResponse := func(response *fern.ListUsersExtendedResponse) *core.PageResponse[*uuid.UUID, *fern.User] {
		next := response.Next
		var results []*fern.User
		if response.Data != nil {
			results = response.Data.Users
		}
		return &core.PageResponse[*uuid.UUID, *fern.User]{
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

func (c *Client) ListUsernames(
	ctx context.Context,
	request *fern.ListUsernamesRequest,
	opts ...option.RequestOption,
) (*core.Page[string], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*string]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("starting_after", fmt.Sprintf("%v", *pageRequest.Cursor))
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
	readPageResponse := func(response *fern.UsernameCursor) *core.PageResponse[*string, string] {
		var next *string
		if response.Cursor != nil {
			next = response.Cursor.After
		}
		var results []string
		if response.Cursor != nil {
			results = response.Cursor.Data
		}
		return &core.PageResponse[*string, string]{
			Next:    next,
			Results: results,
		}
	}
	pager := core.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.StartingAfter)
}

func (c *Client) ListWithGlobalConfig(
	ctx context.Context,
	request *fern.ListWithGlobalConfigRequest,
	opts ...option.RequestOption,
) (*core.Page[string], error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "users"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	prepareCall := func(pageRequest *core.PageRequest[*int]) *core.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("offset", fmt.Sprintf("%v", *pageRequest.Cursor))
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
	next := 1
	if request.Offset != nil {
		next = *request.Offset
	}
	readPageResponse := func(response *fern.UsernameContainer) *core.PageResponse[*int, string] {
		next += 1
		results := response.Results
		return &core.PageResponse[*int, string]{
			Next:    &next,
			Results: results,
		}
	}
	pager := core.NewOffsetPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, &next)
}
