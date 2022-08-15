package statement

// The Verb defines the action between an Actor and an Activity.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#243-verb
type Verb struct {
	ID      string      `json:"id" xapi:"required"`
	Display LanguageMap `json:"display" xapi:"required"`
}

// Creates a new verb
func NewVerb(id string, display LanguageMap) *Verb {
	return &Verb{
		ID:      id,
		Display: display,
	}
}
