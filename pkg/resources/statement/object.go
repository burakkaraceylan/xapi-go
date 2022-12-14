package statement

import (
	"encoding/json"
	"errors"
)

// Actor interface. Impelemented by Agent, Group, StatementRef and SubStatement types.
type IObject interface {
	GetObjectType() string
}

// UnmarshalObject is used to unmarshal an IObject into its type
func UnmarshalObject(m json.RawMessage, object *IObject) error {
	var obj struct {
		ObjectType string `json:"objectType"`
	}

	if err := json.Unmarshal(m, &obj); err != nil {
		return err
	}

	switch obj.ObjectType {
	case "Agent":
		*object = new(Agent)
	case "Group":
		*object = new(Group)
	case "Activity":
		*object = new(Activity)
	case "StatementRef":
		*object = new(StatementRef)
	case "SubStatement":
		*object = new(SubStatement)
	default:
		return errors.New("unknown object type")
	}

	return json.Unmarshal(m, object)
}
