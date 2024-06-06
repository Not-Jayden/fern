// This file was auto-generated by Fern from our API Definition.

package queryparameters

import (
	json "encoding/json"
	fmt "fmt"
	uuid "github.com/google/uuid"
	core "github.com/query-parameters/fern/core"
	time "time"
)

type GetUsersRequest struct {
	Limit          *int              `json:"-" url:"limit,omitempty"`
	Id             *uuid.UUID        `json:"-" url:"id,omitempty"`
	Date           *time.Time        `json:"-" url:"date,omitempty" format:"date"`
	Deadline       *time.Time        `json:"-" url:"deadline,omitempty"`
	Bytes          *[]byte           `json:"-" url:"bytes,omitempty"`
	User           *User             `json:"-" url:"user,omitempty"`
	KeyValue       map[string]string `json:"-" url:"keyValue,omitempty"`
	OptionalString *string           `json:"-" url:"optionalString,omitempty"`
	NestedUser     *NestedUser       `json:"-" url:"nestedUser,omitempty"`
	OptionalUser   *User             `json:"-" url:"optionalUser,omitempty"`
	ExcludeUser    []*User           `json:"-" url:"excludeUser,omitempty"`
	Filter         []*string         `json:"-" url:"filter,omitempty"`
}

type NestedUser struct {
	Name string `json:"name" url:"name"`
	User *User  `json:"user,omitempty" url:"user,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (n *NestedUser) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NestedUser) UnmarshalJSON(data []byte) error {
	type unmarshaler NestedUser
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NestedUser(value)

	extraProperties, err := core.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties

	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *NestedUser) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type User struct {
	Name string   `json:"name" url:"name"`
	Tags []string `json:"tags,omitempty" url:"tags,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *User) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *User) UnmarshalJSON(data []byte) error {
	type unmarshaler User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = User(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) String() string {
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
