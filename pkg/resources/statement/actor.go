package statement

import (
	"encoding/json"
	"errors"
)

// Actor interface. Impelemented by Agent and Group types.
type IActor interface {
	GetActorType() string
}

// UnmarshalActor is used to unmarshal an IActor into its type
func UnmarshalActor(m json.RawMessage, actor *IActor) error {
	var obj struct {
		ObjectType string `json:"objectType"`
	}

	if err := json.Unmarshal(m, &obj); err != nil {
		return err
	}

	switch obj.ObjectType {
	case "Agent":
		*actor = new(Agent)
	case "Group":
		*actor = new(Group)
	default:
		return errors.New("unknown actor type")
	}

	return json.Unmarshal(m, actor)
}
