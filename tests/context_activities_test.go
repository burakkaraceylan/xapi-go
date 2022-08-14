package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ContextActivitiesTestSuite struct {
	suite.Suite
}

func (suite *ContextActivitiesTestSuite) TestEmpty() {
	a := statement.NewContextActivityList()

	assert.IsType(suite.T(), &statement.ContextActivities{}, a)
	assert.IsType(suite.T(), []statement.Activity{}, a.Category)
	assert.IsType(suite.T(), []statement.Activity{}, a.Grouping)
	assert.IsType(suite.T(), []statement.Activity{}, a.Parent)
	assert.IsType(suite.T(), []statement.Activity{}, a.Other)

	assert.Nil(suite.T(), a.Category)
	assert.Nil(suite.T(), a.Grouping)
	assert.Nil(suite.T(), a.Parent)
	assert.Nil(suite.T(), a.Other)
}

func (suite *ContextActivitiesTestSuite) TestInitalize() {
	a := statement.NewContextActivityList()

	cat := statement.NewActivityWithDefiniton("http://github.com/bkaraceylan/xapi-go/test/0", &statement.ActivityDefinition{
		Type: utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/test"),
	})

	a.Append("Category", *cat)
	a.Append("Category", *cat)

	grp := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-group")

	a.Append("Grouping", *grp)
	a.Append("Grouping", *grp)

	parent := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-parent")

	a.Append("Parent", *parent)
	a.Append("Parent", *parent)

	other := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-other")

	a.Append("Other", *other)
	a.Append("Other", *other)

	assert.Equal(suite.T(), 2, len(a.Category))
	assert.Equal(suite.T(), 2, len(a.Grouping))
	assert.Equal(suite.T(), 2, len(a.Parent))
	assert.Equal(suite.T(), 2, len(a.Other))
}

func (suite *ContextActivitiesTestSuite) TestInvalidType() {
	a := statement.NewContextActivityList()

	cat := statement.NewActivityWithDefiniton("http://github.com/bkaraceylan/xapi-go/test/0", &statement.ActivityDefinition{
		Type: utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/test"),
	})

	a.Append("Category", *cat)
	a.Append("Wrong", *cat)

	grp := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-group")

	a.Append("Grouping", *grp)
	a.Append("Wrong", *grp)

	parent := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-parent")

	a.Append("Parent", *parent)
	a.Append("Wrong", *parent)

	other := statement.NewActivity("http://github.com/bkaraceylan/xapi-go/test-other")

	a.Append("Other", *other)
	a.Append("Wrong", *other)

	assert.Equal(suite.T(), 1, len(a.Category))
	assert.Equal(suite.T(), 1, len(a.Grouping))
	assert.Equal(suite.T(), 1, len(a.Parent))
	assert.Equal(suite.T(), 1, len(a.Other))
}

func TestContextActivitiesTestSuite(t *testing.T) {
	suite.Run(t, new(ContextActivitiesTestSuite))
}
