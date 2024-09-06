// This file was auto-generated by Fern from our API Definition.

package teammembers

import (
	context "context"
	fern "github.com/mixed-file-directory/fern"
	core "github.com/mixed-file-directory/fern/core"
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

// Creates a single `TeamMember` object. The `TeamMember` object is returned on successful creates.
// You must provide the following values in your request to this endpoint:
//
// - `given_name`
// - `family_name`
//
// Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#createteammember).
func (c *Client) Create(
	ctx context.Context,
	request *fern.CreateTeamMemberRequest,
	opts ...option.RequestOption,
) (*fern.CreateTeamMemberResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/team-members"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.CreateTeamMemberResponse
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

// Creates multiple `TeamMember` objects. The created `TeamMember` objects are returned on successful creates.
// This process is non-transactional and processes as much of the request as possible. If one of the creates in
// the request cannot be successfully processed, the request is not marked as failed, but the body of the response
// contains explicit error information for the failed create.
//
// Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-create-team-members).
func (c *Client) BatchCreate(
	ctx context.Context,
	request *fern.BatchCreateTeamMembersRequest,
	opts ...option.RequestOption,
) (*fern.BatchCreateTeamMembersResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/team-members/bulk-create"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.BatchCreateTeamMembersResponse
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

// Updates multiple `TeamMember` objects. The updated `TeamMember` objects are returned on successful updates.
// This process is non-transactional and processes as much of the request as possible. If one of the updates in
// the request cannot be successfully processed, the request is not marked as failed, but the body of the response
// contains explicit error information for the failed update.
// Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-update-team-members).
func (c *Client) BatchUpdate(
	ctx context.Context,
	request *fern.BatchUpdateTeamMembersRequest,
	opts ...option.RequestOption,
) (*fern.BatchUpdateTeamMembersResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/team-members/bulk-update"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.BatchUpdateTeamMembersResponse
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

// Returns a paginated list of `TeamMember` objects for a business.
// The list can be filtered by the following:
//
// - location IDs
// - `status`
func (c *Client) Search(
	ctx context.Context,
	request *fern.SearchTeamMembersRequest,
	opts ...option.RequestOption,
) (*fern.SearchTeamMembersResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v2/team-members/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.SearchTeamMembersResponse
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

// Retrieves a `TeamMember` object for the given `TeamMember.id`.
// Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrieve-a-team-member).
func (c *Client) Get(
	ctx context.Context,
	// The ID of the team member to retrieve.
	teamMemberId string,
	opts ...option.RequestOption,
) (*fern.GetTeamMemberResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/team-members/%v", teamMemberId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.GetTeamMemberResponse
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

// Updates a single `TeamMember` object. The `TeamMember` object is returned on successful updates.
// Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#update-a-team-member).
func (c *Client) Update(
	ctx context.Context,
	// The ID of the team member to update.
	teamMemberId string,
	request *fern.UpdateTeamMemberRequest,
	opts ...option.RequestOption,
) (*fern.UpdateTeamMemberResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://connect.squareupsandbox.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v2/team-members/%v", teamMemberId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *fern.UpdateTeamMemberResponse
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
