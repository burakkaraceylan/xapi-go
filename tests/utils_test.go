package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *UtilsTestSuite) TestToJSON() {
	actor := statement.Agent{
		ObjectType: "Agent",
		Name:       utils.Ptr("Foo Bar"),
		Mbox:       utils.Ptr("mailto:foo@bar.com"),
	}

	verb := statement.Verb{
		ID:      "http://adlnet.gov/expapi/verbs/initialized",
		Display: statement.LanguageMap{"en-US": "initialized"},
	}

	statement.NewActivityWithDefiniton("http://id.tincanapi.com/activity/tincan-prototypes/tetris", &statement.ActivityDefinition{
		Name:        &statement.LanguageMap{"en-US": "Js Tetris - Tin Can Prototype"},
		Description: &statement.LanguageMap{"en-US": "A game of tetris."},
		Type:        utils.Ptr("http://activitystrea.ms/schema/1.0/game"),
	})

	object := statement.NewActivityWithDefiniton("http://id.tincanapi.com/activity/tincan-prototypes/tetris", &statement.ActivityDefinition{
		Name:        &statement.LanguageMap{"en-US": "Js Tetris - Tin Can Prototype"},
		Description: &statement.LanguageMap{"en-US": "A game of tetris."},
		Type:        utils.Ptr("http://activitystrea.ms/schema/1.0/game"),
	})

	cat1 := statement.NewActivityWithDefiniton("http://id.tincanapi.com/recipe/tincan-prototypes/tetris/1", &statement.ActivityDefinition{
		Type: utils.Ptr("http://id.tincanapi.com/activitytype/recipe"),
	})

	cat2 := statement.NewActivityWithDefiniton("http://id.tincanapi.com/activity/tincan-prototypes/tetris-template", &statement.ActivityDefinition{
		Type: utils.Ptr("http://id.tincanapi.com/activitytype/source"),
	})

	grp1 := statement.NewActivity("http://id.tincanapi.com/activity/tincan-prototypes")

	grp2 := statement.NewActivity("http://id.tincanapi.com/activity/tincan-prototypes/tetris")

	activities := statement.NewContextActivityList()
	activities.Append("Category", *cat1)
	activities.Append("Category", *cat2)
	activities.Append("Grouping", *grp1)
	activities.Append("Grouping", *grp2)

	context := statement.Context{
		Registration:      utils.Ptr("e168d6a3-46b2-4233-82e7-66b73a179727"),
		ContextActivities: activities,
	}

	authority := *statement.NewAnonymousAgentWithAccount(&statement.Account{
		Name:     "anonymous",
		HomePage: "http://cloud.scorm.com",
	})

	opts := statement.StatementOptions{
		Context:   &context,
		Version:   utils.Ptr("1.0.0"),
		Authority: &authority,
	}

	stmt1 := *statement.NewStatement(actor, verb, object, &opts)

	s, err := utils.ToJson(stmt1, false)

	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(s), 0)

	s2, err := utils.ToJson(stmt1, true)

	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(s2), len(s))
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
