package statement

// A user account on an existing system, such as a private system (LMS or intranet) or a public system (social networking site).
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#agentaccount
type Account struct {
	HomePage string `json:"homePage"`
	Name     string `json:"name"`
}

// NewAccount initializes an account struct
func NewAccount(homePage string, name string) *Account {
	return &Account{
		HomePage: homePage,
		Name:     name,
	}
}
