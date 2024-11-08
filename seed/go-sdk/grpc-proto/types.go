// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/grpc-proto/fern/core"
)

type CreateResponse struct {
	User *UserModel `json:"user,omitempty" url:"user,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateResponse) String() string {
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

type Metadata struct {
	StringMetadataValueMap map[string]*MetadataValue
	StringUnknownMap       map[string]interface{}

	typ string
}

func NewMetadataFromStringMetadataValueMap(value map[string]*MetadataValue) *Metadata {
	return &Metadata{typ: "StringMetadataValueMap", StringMetadataValueMap: value}
}

func NewMetadataFromStringUnknownMap(value map[string]interface{}) *Metadata {
	return &Metadata{typ: "StringUnknownMap", StringUnknownMap: value}
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	var valueStringMetadataValueMap map[string]*MetadataValue
	if err := json.Unmarshal(data, &valueStringMetadataValueMap); err == nil {
		m.typ = "StringMetadataValueMap"
		m.StringMetadataValueMap = valueStringMetadataValueMap
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		m.typ = "StringUnknownMap"
		m.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	if m.typ == "StringMetadataValueMap" || m.StringMetadataValueMap != nil {
		return json.Marshal(m.StringMetadataValueMap)
	}
	if m.typ == "StringUnknownMap" || m.StringUnknownMap != nil {
		return json.Marshal(m.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MetadataVisitor interface {
	VisitStringMetadataValueMap(map[string]*MetadataValue) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (m *Metadata) Accept(visitor MetadataVisitor) error {
	if m.typ == "StringMetadataValueMap" || m.StringMetadataValueMap != nil {
		return visitor.VisitStringMetadataValueMap(m.StringMetadataValueMap)
	}
	if m.typ == "StringUnknownMap" || m.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(m.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MetadataValue struct {
	Double  float64
	String  string
	Boolean bool

	typ string
}

func NewMetadataValueFromDouble(value float64) *MetadataValue {
	return &MetadataValue{typ: "Double", Double: value}
}

func NewMetadataValueFromString(value string) *MetadataValue {
	return &MetadataValue{typ: "String", String: value}
}

func NewMetadataValueFromBoolean(value bool) *MetadataValue {
	return &MetadataValue{typ: "Boolean", Boolean: value}
}

func (m *MetadataValue) UnmarshalJSON(data []byte) error {
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		m.typ = "Double"
		m.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		m.typ = "Boolean"
		m.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MetadataValue) MarshalJSON() ([]byte, error) {
	if m.typ == "Double" || m.Double != 0 {
		return json.Marshal(m.Double)
	}
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return json.Marshal(m.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MetadataValueVisitor interface {
	VisitDouble(float64) error
	VisitString(string) error
	VisitBoolean(bool) error
}

func (m *MetadataValue) Accept(visitor MetadataValueVisitor) error {
	if m.typ == "Double" || m.Double != 0 {
		return visitor.VisitDouble(m.Double)
	}
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return visitor.VisitBoolean(m.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type UserModel struct {
	Username *string   `json:"username,omitempty" url:"username,omitempty"`
	Email    *string   `json:"email,omitempty" url:"email,omitempty"`
	Age      *int      `json:"age,omitempty" url:"age,omitempty"`
	Weight   *float64  `json:"weight,omitempty" url:"weight,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty" url:"metadata,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UserModel) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UserModel) UnmarshalJSON(data []byte) error {
	type unmarshaler UserModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserModel(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserModel) String() string {
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
