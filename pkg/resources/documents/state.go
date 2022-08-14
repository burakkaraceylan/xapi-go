package documents

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement"

type StateDocument struct {
	Document
	Activity     statement.Activity `json:"activity"`
	Agent        statement.Agent    `json:"agent"`
	Registration *string            `json:"registration"`
}
