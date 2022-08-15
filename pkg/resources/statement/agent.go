package statement

import "encoding/json"

// The Actor defines who performed the action. The Actor of a Statement can be an Agent or a Group.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#actor
type Agent struct {
	ObjectType  string   `json:"objectType" xapi:"required"`
	Name        *string  `json:"name,omitempty" xapi:"optional"`
	Mbox        *string  `json:"mbox,omitempty" xapi:"optional"`
	MboxSHA1Sum *string  `json:"mbox_sha1sum,omitempty" xapi:"optional"`
	OpenID      *string  `json:"open_id,omitempty" xapi:"optional"`
	Account     *Account `json:"account,omitempty" xapi:"optional"`
}

// Returns Agent
func (a Agent) GetObjectType() string {
	return "Agent"
}

// Returns Agent
func (a Agent) GetActorType() string {
	return "Agent"
}

// !we ignore possible erros here
// Helper func to marshal into json
func (a *Agent) ToJSON() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// Creates a new agent with a mailbox
func NewAgentWithMbox(name string, mbox string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		Mbox:       &mbox,
	}
}

// Creates a new agent with a mail SHA1
func NewAgentWithSHA1(name string, sha1 string) *Agent {
	return &Agent{
		ObjectType:  "Agent",
		Name:        &name,
		MboxSHA1Sum: &sha1,
	}
}

// Creates a new agent with an openid
func NewAgentWithOpenID(name string, openid string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		OpenID:     &openid,
	}
}

// Creates a new agent with an account
func NewAgentWithAccount(name string, account *Account) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		Account:    account,
	}
}

// Creates a new anonymous agent with a mailbox
func NewAnonymousAgentWithMbox(mbox string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Mbox:       &mbox,
	}
}

// Creates a new anonymous agent with a mail SHA1
func NewAnonymousAgentWithSHA1(sha1 string) *Agent {
	return &Agent{
		ObjectType:  "Agent",
		MboxSHA1Sum: &sha1,
	}
}

// Creates a new anonymous agent with an openid
func NewAnonymousAgentWithOpenID(openid string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		OpenID:     &openid,
	}
}

// Creates a new anonymous agent with an account
func NewAnonymousAgentWithAccount(account *Account) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Account:    account,
	}
}
