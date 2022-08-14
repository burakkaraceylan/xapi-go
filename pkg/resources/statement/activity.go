package statement

import "github.com/burakkaraceylan/xapi-go/pkg/utils"

// The Object defines the thing that was acted on. The Object of a Statement can be an Activity, Agent/Group, SubStatement, or Statement
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#244-object
type Activity struct {
	ID         string              `json:"id" xApi:"required"`
	ObjectType *string             `json:"objectType,omitempty" xApi:"optional"`
	Definition *ActivityDefinition `json:"definition,omitempty" xApi:"optional"`
}

func (a Activity) GetObjectType() string {
	return "Activity"
}

func NewActivity(id string) *Activity {
	return &Activity{
		ID:         id,
		ObjectType: utils.Ptr("Activity"),
	}
}

func NewActivityWithDefiniton(id string, definition *ActivityDefinition) *Activity {
	return &Activity{
		ID:         id,
		ObjectType: utils.Ptr("Activity"),
		Definition: definition,
	}
}
