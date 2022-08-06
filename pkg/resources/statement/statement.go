package statement

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

// Statement represents an evidence for any sort of experience or event which is to be tracked in xAPI.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#statements
type Statement struct {
	ID          *string                  `json:"id,omitempty" xapi:"recommended"`
	Actor       properties.Actor         `json:"actor" xapi:"required"`
	Verb        properties.Verb          `json:"verb" xapi:"required"`
	Object      properties.Object        `json:"object" xapi:"required"`
	Result      *properties.Result       `json:"result,omitempty"  xapi:"optional"`
	Context     *properties.Context      `json:"context,omitempty"  xapi:"optional"`
	Timestamp   *time.Time               `json:"timestamp,omitempty"  xapi:"optional"`
	Stored      *time.Time               `json:"stored,omitempty"  xapi:"optional"`
	Authority   *properties.Actor        `json:"authority,omitempty" xapi:"optional"`
	Version     *string                  `json:"version,omitempty" xapi:"optional"`
	Attachments *[]properties.Attachment `json:"attachments,omitempty" xapi:"optional"`
}

// Outputs the statement as Json
func (stmt *Statement) ToJson(pretty bool) (string, error) {
	var jsonr string
	b, err := json.Marshal(stmt)

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

type MoreStatements struct {
	More       string      `json:"more"`
	Statements []Statement `json:"statements"`
}
