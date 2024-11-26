// This file was auto-generated by Fern from our API Definition.

package folderd

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/audiences/fern/internal"
)

type Response struct {
	Foo string `json:"foo" url:"foo"`

	extraProperties map[string]interface{}
}

func (r *Response) GetFoo() string {
	if r == nil {
		return ""
	}
	return r.Foo
}

func (r *Response) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *Response) UnmarshalJSON(data []byte) error {
	type unmarshaler Response
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = Response(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *Response) String() string {
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}
