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

func (a Agent) GetObjectType() string {
	return "Agent"
}

func (a Agent) GetActorType() string {
	return "Agent"
}

// !we ignore possible erros here
func (a *Agent) ToJSON() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func NewAgentWithMbox(name string, mbox string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		Mbox:       &mbox,
	}
}

func NewAgentWithSHA1(name string, sha1 string) *Agent {
	return &Agent{
		ObjectType:  "Agent",
		Name:        &name,
		MboxSHA1Sum: &sha1,
	}
}

func NewAgentWithOpenID(name string, openid string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		OpenID:     &openid,
	}
}

func NewAgentWithAccount(name string, account *Account) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Name:       &name,
		Account:    account,
	}
}

func NewAnonymousAgentWithMbox(mbox string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Mbox:       &mbox,
	}
}

func NewAnonymousAgentWithSHA1(sha1 string) *Agent {
	return &Agent{
		ObjectType:  "Agent",
		MboxSHA1Sum: &sha1,
	}
}

func NewAnonymousAgentWithOpenID(openid string) *Agent {
	return &Agent{
		ObjectType: "Agent",
		OpenID:     &openid,
	}
}

func NewAnonymousAgentWithAccount(account *Account) *Agent {
	return &Agent{
		ObjectType: "Agent",
		Account:    account,
	}
}
