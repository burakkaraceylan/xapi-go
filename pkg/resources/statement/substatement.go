package statement

import (
	"time"
)

type SubStatement struct {
	Statement
	ObjectType string `json:"objectType" xapi:"required"`
}

func (s SubStatement) GetObjectType() string {
	return "SubStatement"
}

type SubStatementOptions struct {
	Result      *Result      `json:"result,omitempty"  xapi:"optional"`
	Context     *Context     `json:"context,omitempty"  xapi:"optional"`
	Timestamp   *time.Time   `json:"timestamp,omitempty"  xapi:"optional"`
	Attachments []Attachment `json:"attachments,omitempty" xapi:"optional"`
}

func NewSubStatement(actor IActor, verb Verb, object IObject, options ...*SubStatementOptions) *SubStatement {
	statement := new(Statement)

	statement.Actor = actor
	statement.Verb = verb
	statement.Object = object

	if len(options) > 0 {
		statement.Result = options[0].Result
		statement.Context = options[0].Context
		statement.Timestamp = options[0].Timestamp
		statement.Attachments = options[0].Attachments
	}

	return &SubStatement{
		Statement:  *statement,
		ObjectType: "SubStatement",
	}
}
