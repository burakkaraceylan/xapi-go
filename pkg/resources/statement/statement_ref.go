package statement

type StatementRef struct {
	ID         string `json:"id" xapi:"required"`
	ObjectType string `json:"objectType" xapi:"required"`
}

func (s StatementRef) GetObjectType() string {
	return "StatementRef"
}

func NewStatementRef(id string) *StatementRef {
	return &StatementRef{
		ID:         id,
		ObjectType: "StatementRef",
	}
}
