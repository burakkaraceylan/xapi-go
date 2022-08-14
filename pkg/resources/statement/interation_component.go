package statement

type InteractionComponent struct {
	ID          string       `json:"id" xapi:"required"`
	Description *LanguageMap `json:"description" xapi:"optional"`
}

func NewInteractionComponent(id string, params ...*LanguageMap) *InteractionComponent {
	ic := InteractionComponent{
		ID: id,
	}

	if len(params) == 1 && params[0] != nil {
		ic.Description = params[0]
	}

	return &ic
}
