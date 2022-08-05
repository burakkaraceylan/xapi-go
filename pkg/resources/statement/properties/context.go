package properties

// An optional property that provides a place to add contextual information to a Statement. All "context" properties are optional.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#246-context
type Context struct {
	Registration      *string            `json:"registration,omitempty" xapi:"optional"`
	Instructor        *Actor             `json:"agent,omitempty" xapi:"optional"`
	Team              *Actor             `json:"team,omitempty" xapi:"optional"`
	ContextActivities *ContextActivities `json:"contextActivities,omitempty" xapi:"optional"`
	Revision          *string            `json:"revision,omitempty" xapi:"optional"`
	Platform          *string            `json:"platform,omitempty" xapi:"optional"`
	Language          *string            `json:"language,omitempty" xapi:"optional"`
	Statement         *Object            `json:"statement,omitempty" xapi:"optional"`
	Extensions        *Extensions        `json:"extensions,omitempty" xapi:"optional"`
}
