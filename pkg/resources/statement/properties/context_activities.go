package properties

// A map of the types of learning activity context that this Statement is related to.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#246-context
type ContextActivities struct {
	Parent   []Object `json:"parent"`
	Grouping []Object `json:"grouping"`
	Category []Object `json:"category"`
	Other    []Object `json:"other"`
}
