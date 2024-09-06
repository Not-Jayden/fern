// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mixed-file-directory/fern/core"
)

type CreateMobileAuthorizationCodeRequest struct {
	// The Square location ID that the authorization code should be tied to.
	LocationId *string `json:"location_id,omitempty" url:"-"`
}

// Defines the fields that are included in the response body of
// a request to the `CreateMobileAuthorizationCode` endpoint.
type CreateMobileAuthorizationCodeResponse struct {
	// The generated authorization code that connects a mobile application instance
	// to a Square account.
	AuthorizationCode *string `json:"authorization_code,omitempty" url:"authorization_code,omitempty"`
	// The timestamp when `authorization_code` expires, in
	// [RFC 3339](https://tools.ietf.org/html/rfc3339) format (for example, "2016-09-04T23:59:33.123Z").
	ExpiresAt *string `json:"expires_at,omitempty" url:"expires_at,omitempty"`
	// An error object that provides details about how creation of the authorization
	// code failed.
	Error *Error `json:"error,omitempty" url:"error,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateMobileAuthorizationCodeResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateMobileAuthorizationCodeResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateMobileAuthorizationCodeResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateMobileAuthorizationCodeResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateMobileAuthorizationCodeResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
