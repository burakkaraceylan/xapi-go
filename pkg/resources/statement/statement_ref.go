package statement

// A Statement Reference is a pointer to another pre-existing Statement.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#statement-references
type StatementRef struct {
	ID         string `json:"id" xapi:"required"`
	ObjectType string `json:"objectType" xapi:"required"`
}

// Returns StatementRef
func (s StatementRef) GetObjectType() string {
	return "StatementRef"
}

// Creates a new statementref
func NewStatementRef(id string) *StatementRef {
	return &StatementRef{
		ID:         id,
		ObjectType: "StatementRef",
	}
}
