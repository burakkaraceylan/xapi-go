package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AgentTestSuite struct {
	suite.Suite
}

func (suite *AgentTestSuite) TestInitializeMbox() {
	a := statement.NewAgentWithMbox("Burak Karaceylan", "mailto:bkaraceylan@gmail.com")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "Burak Karaceylan", *a.Name)
	assert.Equal(suite.T(), "mailto:bkaraceylan@gmail.com", *a.Mbox)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.OpenID)

}

func (suite *AgentTestSuite) TestInitializeMboxSha() {
	a := statement.NewAgentWithSHA1("Burak Karaceylan", "a6a8da09cfdc62147ffa37ba5906bf72d90b7902")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "Burak Karaceylan", *a.Name)
	assert.Equal(suite.T(), "a6a8da09cfdc62147ffa37ba5906bf72d90b7902", *a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.OpenID)

}

func (suite *AgentTestSuite) TestInitializeOpenID() {
	a := statement.NewAgentWithOpenID("Burak Karaceylan", "test")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "Burak Karaceylan", *a.Name)
	assert.Equal(suite.T(), "test", *a.OpenID)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.MboxSHA1Sum)

}

func (suite *AgentTestSuite) TestInitializeAccount() {
	acc := statement.NewAccount("http://github.com", "bkaraceylan")
	a := statement.NewAgentWithAccount("Burak Karaceylan", acc)

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "Burak Karaceylan", *a.Name)
	assert.Equal(suite.T(), *acc, *a.Account)
	assert.Nil(suite.T(), a.OpenID)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.MboxSHA1Sum)

}
func (suite *AgentTestSuite) TestInitializeAnonMbox() {
	a := statement.NewAnonymousAgentWithMbox("mailto:bkaraceylan@gmail.com")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "mailto:bkaraceylan@gmail.com", *a.Mbox)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.OpenID)
	assert.Nil(suite.T(), a.Name)

}

func (suite *AgentTestSuite) TestInitializeAnonMboxSha() {
	a := statement.NewAnonymousAgentWithSHA1("a6a8da09cfdc62147ffa37ba5906bf72d90b7902")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "a6a8da09cfdc62147ffa37ba5906bf72d90b7902", *a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.OpenID)
	assert.Nil(suite.T(), a.Name)
}

func (suite *AgentTestSuite) TestInitializeAnonOpenID() {
	a := statement.NewAnonymousAgentWithOpenID("test")

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), "test", *a.OpenID)
	assert.Nil(suite.T(), a.Account)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.Name)
}

func (suite *AgentTestSuite) TestInitializeAnonAccount() {
	acc := statement.NewAccount("http://github.com", "bkaraceylan")
	a := statement.NewAnonymousAgentWithAccount(acc)

	assert.IsType(suite.T(), &statement.Agent{}, a)
	assert.Equal(suite.T(), "Agent", a.ObjectType)
	assert.Equal(suite.T(), "Agent", a.GetActorType())
	assert.Equal(suite.T(), "Agent", a.GetObjectType())
	assert.Equal(suite.T(), *acc, *a.Account)
	assert.Nil(suite.T(), a.OpenID)
	assert.Nil(suite.T(), a.Mbox)
	assert.Nil(suite.T(), a.MboxSHA1Sum)
	assert.Nil(suite.T(), a.Name)

}

func TestAgentTestSuite(t *testing.T) {
	suite.Run(t, new(AgentTestSuite))
}
