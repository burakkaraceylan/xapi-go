package statement

type Group struct {
	ObjectType string  `json:"objectType" xapi:"required"`
	Name       *string `json:"name,omitempty" xapi:"optional"`
	Members    []Agent `json:"members" xapi:"required"`
}

func (g Group) GetObjectType() string {
	return "Group"
}

func (g Group) GetActorType() string {
	return "Group"
}

func (g *Group) AddMember(agent Agent) {
	g.Members = append(g.Members, agent)
}

func NewGroup(name string) *Group {
	return &Group{
		ObjectType: "Group",
		Name:       &name,
	}
}

func NewAnonymousGroup() *Group {
	return &Group{
		ObjectType: "Group",
	}
}
