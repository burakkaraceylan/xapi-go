package utils

import (
	"bytes"
	"encoding/json"
)

// Returns pointer to a literal
func Ptr[T any](v T) *T {
	return &v
}

// Outputs the statement as Json
func ToJson(a any, pretty bool) (string, error) {
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
