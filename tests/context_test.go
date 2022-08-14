package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContextTestSuite struct {
	suite.Suite
}

func (suite *ContextTestSuite) TestEmpty() {
	c := statement.Context{}

	s := ""
	sp := utils.Ptr(s)

	assert.IsType(suite.T(), sp, c.Registration)
	assert.IsType(suite.T(), sp, c.Revision)
	assert.IsType(suite.T(), sp, c.Platform)
	assert.IsType(suite.T(), sp, c.Language)
	assert.IsType(suite.T(), (statement.IActor)(nil), c.Instructor)
	assert.IsType(suite.T(), &statement.Group{}, c.Team)
	assert.IsType(suite.T(), &statement.StatementRef{}, c.Statement)
	assert.IsType(suite.T(), &statement.Extensions{}, c.Extensions)
	assert.IsType(suite.T(), &statement.ContextActivities{}, c.ContextActivities)

	assert.Nil(suite.T(), c.Registration)
	assert.Nil(suite.T(), c.Revision)
	assert.Nil(suite.T(), c.Platform)
	assert.Nil(suite.T(), c.Registration)
	assert.Nil(suite.T(), c.Instructor)
	assert.Nil(suite.T(), c.Team)
	assert.Nil(suite.T(), c.Statement)
	assert.Nil(suite.T(), c.Extensions)
	assert.Nil(suite.T(), c.ContextActivities)

}

func (suite *ContextTestSuite) TestInitalize() {

	agent := statement.NewAgentWithMbox("Burak Karaceylan", "mailto:bkaraceylan@gmail.com")
	group := statement.NewGroup("Testing Group")
	group.AddMember(*agent)
	sref := statement.NewStatementRef("test")
	ext := statement.Extensions{"ext1": "test1"}

	a := statement.NewContextActivityList()

	cat := statement.NewActivityWithDefiniton("http://github.com/bkaraceylan/xapi-go/test/0", &statement.ActivityDefinition{
		Type: utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/test"),
	})

	a.Append("Category", *cat)

	s := "test"
	sp := utils.Ptr(s)

	c := statement.Context{
		Registration:      sp,
		Instructor:        agent,
		Team:              group,
		ContextActivities: a,
		Revision:          sp,
		Platform:          sp,
		Language:          sp,
		Statement:         sref,
		Extensions:        &ext,
	}

	assert.Equal(suite.T(), "test", *c.Registration)
	assert.Equal(suite.T(), "test", *c.Revision)
	assert.Equal(suite.T(), "test", *c.Platform)
	assert.Equal(suite.T(), "test", *c.Registration)
	assert.Equal(suite.T(), *agent, *c.Instructor.(*statement.Agent))
	assert.Equal(suite.T(), *group, *c.Team)
	assert.Equal(suite.T(), *sref, *c.Statement)
	assert.Equal(suite.T(), ext, *c.Extensions)
	assert.Equal(suite.T(), *a, *c.ContextActivities)
}

func TestContextTestSuite(t *testing.T) {
	suite.Run(t, new(ContextTestSuite))
}
