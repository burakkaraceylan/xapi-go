package documents

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement"

// Represents an gent profile document
type AgentDocument struct {
	Document
	Agent statement.Agent `json:"agent"`
}
