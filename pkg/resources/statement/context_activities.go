package statement

// A map of the types of learning activity context that this Statement is related to.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#246-context
type ContextActivities struct {
	Parent   []Activity `json:"parent,omitempty"`
	Grouping []Activity `json:"grouping,omitempty"`
	Category []Activity `json:"category,omitempty"`
	Other    []Activity `json:"other,omitempty"`
}

// Append is used to add a new activity to the activity list
// ContextType must be one of Parent, Grouping, Category, or Other
func (c *ContextActivities) Append(ctxType string, obj Activity) {
	switch ctxType {
	case "Parent":
		c.Parent = append(c.Parent, obj)
	case "Grouping":
		c.Grouping = append(c.Grouping, obj)
	case "Category":
		c.Category = append(c.Category, obj)
	case "Other":
		c.Other = append(c.Other, obj)
	default:
		return
	}

}

// NewContextActivityList is used to initialize a new context activity list
func NewContextActivityList() *ContextActivities {
	return &ContextActivities{}
}
