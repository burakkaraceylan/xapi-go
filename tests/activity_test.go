package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ActivityTestSuite struct {
	suite.Suite
}

func (suite *ActivityTestSuite) TestInitialize() {
	a := statement.NewActivity("test")

	assert.IsType(suite.T(), &statement.Activity{}, a)
	assert.Equal(suite.T(), "test", a.ID)
	assert.Equal(suite.T(), "Activity", *a.ObjectType)
	assert.Equal(suite.T(), "Activity", a.GetObjectType())
	assert.Nil(suite.T(), a.Definition)

}

func (suite *ActivityTestSuite) TestInitializeWithDefinition() {
	def := statement.ActivityDefinition{
		Name: &statement.LanguageMap{"en-US": "test"},
	}
	a := statement.NewActivityWithDefiniton("test", &def)

	assert.IsType(suite.T(), &statement.Activity{}, a)
	assert.Equal(suite.T(), "test", a.ID)
	assert.Equal(suite.T(), "Activity", *a.ObjectType)
	assert.Equal(suite.T(), "Activity", a.GetObjectType())
	assert.Equal(suite.T(), def, *a.Definition)

}

func TestActivityTestSuite(t *testing.T) {
	suite.Run(t, new(ActivityTestSuite))
}
