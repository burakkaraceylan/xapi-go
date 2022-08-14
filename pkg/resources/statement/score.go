package statement

// An optional property that represents the outcome of a graded Activity achieved by an Agent.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#2451-score
type Score struct {
	Scaled *float32 `json:"scaled,omitempty" xapi:"recommended"`
	Raw    *float32 `json:"raw,omitempty" xapi:"optional"`
	Min    *float32 `json:"min,omitempty" xapi:"optional"`
	Max    *float32 `json:"max,omitempty" xapi:"optional"`
}
