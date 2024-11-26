// This file was auto-generated by Fern from our API Definition.

package extraproperties

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/extra-properties/fern/internal"
)

type User struct {
	Name string `json:"name" url:"name"`

	ExtraProperties map[string]interface{} `json:"-" url:"-"`
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func (u *User) GetExtraProperties() map[string]interface{} {
	return u.ExtraProperties
}

func (u *User) UnmarshalJSON(data []byte) error {
	type embed User
	var unmarshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*u = User(unmarshaler.embed)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.ExtraProperties = extraProperties
	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	type embed User
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return internal.MarshalJSONWithExtraProperties(marshaler, u.ExtraProperties)
}

func (u *User) String() string {
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}
