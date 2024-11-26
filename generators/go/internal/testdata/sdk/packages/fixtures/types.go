// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/internal"
)

type Error struct {
	Message string `json:"message" url:"message"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *Error) GetMessage() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func (e *Error) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Error) UnmarshalJSON(data []byte) error {
	type unmarshaler Error
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Error(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *Error) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type Foo struct {
	Id   string `json:"id" url:"id"`
	Name string `json:"name" url:"name"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *Foo) GetId() string {
	if f == nil {
		return ""
	}
	return f.Id
}

func (f *Foo) GetName() string {
	if f == nil {
		return ""
	}
	return f.Name
}

func (f *Foo) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *Foo) UnmarshalJSON(data []byte) error {
	type unmarshaler Foo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = Foo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *Foo) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}
