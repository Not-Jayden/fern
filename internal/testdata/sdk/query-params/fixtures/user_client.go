// Generated by Fern. Do not edit.

package api

import (
	context "context"
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params/fixtures/core"
	io "io"
	http "net/http"
)

type UserClient interface{}

type getAllUsersEndpoint struct {
	url    string
	client core.HTTPClient
}

func newgetAllUsersEndpoint(url string, client core.HTTPClient) *getAllUsersEndpoint {
	return &getAllUsersEndpoint{
		url:    url,
		client: client,
	}
}

func (g *getAllUsersEndpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (g *getAllUsersEndpoint) Call(ctx context.Context) (string, error) {
	var response string
	if err := core.DoRequest(
		ctx,
		g.client,
		g.url,
		http.MethodGet,
		nil,
		response,
		nil,
		g.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}
