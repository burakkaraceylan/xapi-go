package properties

import "github.com/bkaraceylan/xapi-go/pkg/resources/statement/special"

// The Verb defines the action between an Actor and an Activity.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#243-verb
type Verb struct {
	ID      string `json:"id"`
	Display special.LanguageMap
}
