package documents

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement"

type AgentDocument struct {
	Document
	Agent statement.Agent `json:"agent"`
}
