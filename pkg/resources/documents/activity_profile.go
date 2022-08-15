package documents

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement"

// Represents activity profile document
type ActivityDocument struct {
	Document
	Activity statement.Activity `json:"activity"`
}
