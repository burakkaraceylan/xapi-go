package about

import (
	"bytes"
	"encoding/json"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

// Returns JSON Object containing information about this LRS, including the xAPI version supported.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Communication.md#description-7
type About struct {
	Version    []string               `json:"version" xapi:"required"`
	Extensions *properties.Extensions `json:"extensions" xapi:"optional"`
}

// Outputs the statement as Json
func (a *About) ToJson(pretty bool) (string, error) {
	var jsonr string
	b, err := json.Marshal(a)

	if err != nil {
		return "", err
	}

	if pretty {
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, b, "", "    "); err != nil {
			return "", err
		}

		jsonr = prettyJSON.String()
	} else {
		jsonr = string(b)
	}

	return jsonr, nil
}
