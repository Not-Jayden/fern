// This file was auto-generated by Fern from our API Definition.

package customauth

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/custom-auth/fern/core"
)

type UnauthorizedRequestErrorBody struct {
	Message string `json:"message" url:"message"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UnauthorizedRequestErrorBody) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UnauthorizedRequestErrorBody) UnmarshalJSON(data []byte) error {
	type unmarshaler UnauthorizedRequestErrorBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnauthorizedRequestErrorBody(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnauthorizedRequestErrorBody) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
