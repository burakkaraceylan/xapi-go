package statement

// An optional property that represents a measured outcome related to the Statement in which it is included.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#245-result
type Result struct {
	Score      *Score      `json:"score,omitempty" xapi:"optional"`
	Success    *bool       `json:"success,omitempty" xapi:"optional"`
	Completion *bool       `json:"completion,omitempty" xapi:"optional"`
	Response   *string     `json:"response,omitempty" xapi:"optional"`
	Duration   *string     `json:"duration,omitempty" xapi:"optional"`
	Extensions *Extensions `json:"extensions,omitempty" xapi:"optional"`
}
