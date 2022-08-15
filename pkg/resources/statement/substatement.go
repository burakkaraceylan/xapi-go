package statement

import (
	"time"
)

// A SubStatement is like a StatementRef in that it is included as part of a containing Statement, but unlike a StatementRef,
// it does not represent an event that has occurred. It can be used to describe, for example, a predication of a potential future
// Statement or the behavior a teacher looked for when evaluating a student (without representing the student actually doing that behavior).
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#substatements
type SubStatement struct {
	Statement
	ObjectType string `json:"objectType" xapi:"required"`
}

// Returns SubStatement
func (s SubStatement) GetObjectType() string {
	return "SubStatement"
}

// SubStatement optional parameters
type SubStatementOptions struct {
	Result      *Result      `json:"result,omitempty"  xapi:"optional"`
	Context     *Context     `json:"context,omitempty"  xapi:"optional"`
	Timestamp   *time.Time   `json:"timestamp,omitempty"  xapi:"optional"`
	Attachments []Attachment `json:"attachments,omitempty" xapi:"optional"`
}

// Creates a new substatement
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
