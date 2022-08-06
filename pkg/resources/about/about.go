package about

import (
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

// Returns JSON Object containing information about this LRS, including the xAPI version supported.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Communication.md#description-7
type About struct {
	Version    []string               `json:"version" xapi:"required"`
	Extensions *properties.Extensions `json:"extensions" xapi:"optional"`
}
