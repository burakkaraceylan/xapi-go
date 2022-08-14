package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GroupTestSuite struct {
	suite.Suite
	Agent statement.Agent
}

func (suite *GroupTestSuite) SetupSuite() {
	suite.Agent = *statement.NewAgentWithMbox("Burak Karaceylan", "mailto:bkaraceylan@gmail.com")
}

func (suite *GroupTestSuite) TestNamedGroup() {
	g := statement.NewGroup("test")
	g.AddMember(suite.Agent)
	g.AddMember(suite.Agent)

	assert.IsType(suite.T(), &statement.Group{}, g)
	assert.Equal(suite.T(), "Group", g.ObjectType)
	assert.Equal(suite.T(), "Group", g.GetActorType())
	assert.Equal(suite.T(), "Group", g.GetObjectType())
	assert.Equal(suite.T(), "test", *g.Name)
	assert.Equal(suite.T(), suite.Agent, g.Members[0])
	assert.Equal(suite.T(), suite.Agent, g.Members[1])

}

func (suite *GroupTestSuite) TestAnonymousGroup() {
	g := statement.NewAnonymousGroup()
	g.AddMember(suite.Agent)
	g.AddMember(suite.Agent)

	assert.IsType(suite.T(), &statement.Group{}, g)
	assert.Equal(suite.T(), "Group", g.ObjectType)
	assert.Equal(suite.T(), "Group", g.GetActorType())
	assert.Equal(suite.T(), "Group", g.GetObjectType())
	assert.Nil(suite.T(), g.Name)
	assert.Equal(suite.T(), suite.Agent, g.Members[0])
	assert.Equal(suite.T(), suite.Agent, g.Members[1])

}

func TestGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GroupTestSuite))
}
