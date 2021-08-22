// Code generated by goa v3.4.3, DO NOT EDIT.
//
// inspect HTTP client CLI support package
//
// Command:
// $ goa gen algodexidx/design

package client

import (
	inspect "algodexidx/gen/inspect"
	"encoding/json"
	"fmt"
)

// BuildUnpackPayload builds the payload for the inspect unpack endpoint from
// CLI flags.
func BuildUnpackPayload(inspectUnpackBody string) (*inspect.UnpackPayload, error) {
	var err error
	var body UnpackRequestBody
	{
		err = json.Unmarshal([]byte(inspectUnpackBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"msgpack\": \"Placeat quidem blanditiis dolorem officiis aspernatur.\"\n   }'")
		}
	}
	v := &inspect.UnpackPayload{
		Msgpack: body.Msgpack,
	}

	return v, nil
}
