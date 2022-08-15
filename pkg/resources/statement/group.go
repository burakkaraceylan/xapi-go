package statement

// A Group represents a collection of Agents and can be used in most of the same situations an Agent can be used.
// There are two types of Groups: Anonymous Groups and Identified Groups.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#group
type Group struct {
	ObjectType string  `json:"objectType" xapi:"required"`
	Name       *string `json:"name,omitempty" xapi:"optional"`
	Members    []Agent `json:"members" xapi:"required"`
}

// Returns Group
func (g Group) GetObjectType() string {
	return "Group"
}

// Returns Group
func (g Group) GetActorType() string {
	return "Group"
}

// Adds a new member to the group
func (g *Group) AddMember(agent Agent) {
	g.Members = append(g.Members, agent)
}

// Creates a new group
func NewGroup(name string) *Group {
	return &Group{
		ObjectType: "Group",
		Name:       &name,
	}
}

// Creates a new anonymous group
func NewAnonymousGroup() *Group {
	return &Group{
		ObjectType: "Group",
	}
}
