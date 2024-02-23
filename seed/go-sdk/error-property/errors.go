// This file was auto-generated by Fern from our API Definition.

package errorproperty

import (
	json "encoding/json"
	core "github.com/error-property/fern/core"
)

type PropertyBasedErrorTest struct {
	*core.APIError
	Body *PropertyBasedErrorTestBody
}

func (p *PropertyBasedErrorTest) UnmarshalJSON(data []byte) error {
	var body *PropertyBasedErrorTestBody
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	p.StatusCode = 400
	p.Body = body
	return nil
}

func (p *PropertyBasedErrorTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Body)
}

func (p *PropertyBasedErrorTest) Unwrap() error {
	return p.APIError
}
