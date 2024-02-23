// This file was auto-generated by Fern from our API Definition.

package objectswithimports

import (
	json "encoding/json"
	fmt "fmt"
	commons "github.com/objects-with-imports/fern/commons"
	core "github.com/objects-with-imports/fern/core"
)

type Node struct {
	Id       string            `json:"id" url:"id"`
	Label    *string           `json:"label,omitempty" url:"label,omitempty"`
	Metadata *commons.Metadata `json:"metadata,omitempty" url:"metadata,omitempty"`

	_rawJSON json.RawMessage
}

func (n *Node) UnmarshalJSON(data []byte) error {
	type unmarshaler Node
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = Node(value)
	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *Node) String() string {
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

type Tree struct {
	Nodes []*Node `json:"nodes,omitempty" url:"nodes,omitempty"`

	_rawJSON json.RawMessage
}

func (t *Tree) UnmarshalJSON(data []byte) error {
	type unmarshaler Tree
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Tree(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *Tree) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type File struct {
	Name     string   `json:"name" url:"name"`
	Contents string   `json:"contents" url:"contents"`
	Info     FileInfo `json:"info,omitempty" url:"info,omitempty"`

	_rawJSON json.RawMessage
}

func (f *File) UnmarshalJSON(data []byte) error {
	type unmarshaler File
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = File(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *File) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FileInfo string

const (
	// A regular file (e.g. foo.txt).
	FileInfoRegular FileInfo = "REGULAR"
	// A directory (e.g. foo/).
	FileInfoDirectory FileInfo = "DIRECTORY"
)

func NewFileInfoFromString(s string) (FileInfo, error) {
	switch s {
	case "REGULAR":
		return FileInfoRegular, nil
	case "DIRECTORY":
		return FileInfoDirectory, nil
	}
	var t FileInfo
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (f FileInfo) Ptr() *FileInfo {
	return &f
}
