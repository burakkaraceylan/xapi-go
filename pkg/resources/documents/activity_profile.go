package documents

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement"

type ActivityDocument struct {
	Document
	Activity statement.Activity `json:"activity"`
}
