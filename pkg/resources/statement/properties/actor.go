package properties

// The Actor defines who performed the action. The Actor of a Statement can be an Agent or a Group.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#actor
type Actor struct {
	ObjectType  string   `json:"objectType" xapi:"required"`
	Name        *string  `json:"name,omitempty" xapi:"optional"`
	Member      *[]Actor `json:"member,omitempty" xapi:"optional"`
	Mbox        *string  `json:"mbox,omitempty" xapi:"optional"`
	MboxSHA1Sum *string  `json:"mbox_sha1sum,omitempty" xapi:"optional"`
	OpenID      *string  `json:"open_id,omitempty" xapi:"optional"`
	Account     *Account `json:"account,omitempty" xapi:"optional"`
}
