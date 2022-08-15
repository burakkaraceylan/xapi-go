package statement

// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#interaction-components
type InteractionComponent struct {
	ID          string       `json:"id" xapi:"required"`
	Description *LanguageMap `json:"description" xapi:"optional"`
}

// Creates a new interaction component
func NewInteractionComponent(id string, params ...*LanguageMap) *InteractionComponent {
	ic := InteractionComponent{
		ID: id,
	}

	if len(params) == 1 && params[0] != nil {
		ic.Description = params[0]
	}

	return &ic
}
