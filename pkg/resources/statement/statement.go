package statement

import (
	"encoding/json"
	"time"
)

// Statement represents an evidence for any sort of experience or event which is to be tracked in xAPI.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#statements
type Statement struct {
	StatementOptions
	Actor  IActor  `json:"actor" xapi:"required"`
	Verb   Verb    `json:"verb" xapi:"required"`
	Object IObject `json:"object" xapi:"required"`
}

// StatementOptions struct contains optional statement parameters
type StatementOptions struct {
	ID          *string      `json:"id,omitempty" xapi:"recommended"`
	Result      *Result      `json:"result,omitempty"  xapi:"optional"`
	Context     *Context     `json:"context,omitempty"  xapi:"optional"`
	Timestamp   *time.Time   `json:"timestamp,omitempty"  xapi:"optional"`
	Stored      *time.Time   `json:"stored,omitempty"  xapi:"optional"`
	Authority   IActor       `json:"authority,omitempty" xapi:"optional"`
	Version     *string      `json:"version,omitempty" xapi:"optional"`
	Attachments []Attachment `json:"attachments,omitempty" xapi:"optional"`
}

// NewStatement produces a Statement sturct.
func NewStatement(actor IActor, verb Verb, object IObject, opt ...*StatementOptions) *Statement {
	statement := new(Statement)

	statement.Actor = actor
	statement.Verb = verb
	statement.Object = object

	if len(opt) > 0 {
		statement.ID = opt[0].ID
		statement.Result = opt[0].Result
		statement.Context = opt[0].Context
		statement.Timestamp = opt[0].Timestamp
		statement.Stored = opt[0].Stored
		statement.Authority = opt[0].Authority
		statement.Version = opt[0].Version
		statement.Attachments = opt[0].Attachments
	}

	return statement
}

// A collection of Statements can be retrieved by performing a query on the Statement Resource
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#retrieval
type StatementResult struct {
	More       string      `json:"more"`
	Statements []Statement `json:"statements"`
}

// Unmarshals the statement. A custom unmarshaller is required due to Actor, Object and Authority fields being interfaces.
func (s *Statement) UnmarshalJSON(data []byte) error {
	raw := struct {
		Actor       json.RawMessage `json:"actor"`
		Verb        Verb            `json:"verb"`
		Object      json.RawMessage `json:"object"`
		ID          *string         `json:"id,omitempty"`
		Result      *Result         `json:"result,omitempty"`
		Context     *Context        `json:"context,omitempty"`
		Timestamp   *time.Time      `json:"timestamp,omitempty"`
		Stored      *time.Time      `json:"stored,omitempty"`
		Authority   json.RawMessage `json:"authority,omitempty"`
		Version     *string         `json:"version,omitempty"`
		Attachments []Attachment    `json:"attachments,omitempty"`
	}{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	s.Verb = raw.Verb
	s.ID = raw.ID
	s.Result = raw.Result
	s.Context = raw.Context
	s.Timestamp = raw.Timestamp
	s.Stored = raw.Stored
	s.Version = raw.Version
	s.Attachments = raw.Attachments

	if err = UnmarshalActor(raw.Actor, &s.Actor); err != nil {
		return err
	}

	if err = UnmarshalObject(raw.Object, &s.Object); err != nil {
		return err
	}

	if err = UnmarshalActor(raw.Authority, &s.Authority); err != nil {
		return err
	}

	return nil
}
