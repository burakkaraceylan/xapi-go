package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (suite *ResourceTestSuite) TestToJSON() {
	actor := properties.Actor{
		ObjectType: "Agent",
		Name:       utils.Ptr("Foo Bar"),
		Mbox:       utils.Ptr("mailto:foo@bar.com"),
	}

	verb := properties.Verb{
		ID:      "http://adlnet.gov/expapi/verbs/initialized",
		Display: special.LanguageMap{"en-US": "initialized"},
	}

	object := properties.Object{
		ID:         "http://id.tincanapi.com/activity/tincan-prototypes/tetris",
		ObjectType: utils.Ptr("Activity"),
		Definition: &properties.Definition{
			Name:        &special.LanguageMap{"en-US": "Js Tetris - Tin Can Prototype"},
			Description: &special.LanguageMap{"en-US": "A game of tetris."},
			Type:        utils.Ptr("http://activitystrea.ms/schema/1.0/game"),
		},
	}

	cat1 := properties.Object{
		ID:         "http://id.tincanapi.com/recipe/tincan-prototypes/tetris/1",
		ObjectType: utils.Ptr("Activity"),
		Definition: &properties.Definition{
			Type: utils.Ptr("http://id.tincanapi.com/activitytype/recipe"),
		},
	}

	cat2 := properties.Object{
		ID:         "http://id.tincanapi.com/activity/tincan-prototypes/tetris-template",
		ObjectType: utils.Ptr("Activity"),
		Definition: &properties.Definition{
			Type: utils.Ptr("http://id.tincanapi.com/activitytype/source"),
		},
	}

	grp1 := properties.Object{
		ID:         "http://id.tincanapi.com/activity/tincan-prototypes",
		ObjectType: utils.Ptr("Activity"),
	}

	grp2 := properties.Object{
		ID:         "http://id.tincanapi.com/activity/tincan-prototypes/tetris",
		ObjectType: utils.Ptr("Activity"),
	}

	context := properties.Context{
		Registration: utils.Ptr("e168d6a3-46b2-4233-82e7-66b73a179727"),
		ContextActivities: &properties.ContextActivities{
			Category: &[]properties.Object{cat1, cat2},
			Grouping: &[]properties.Object{grp1, grp2},
		},
	}

	authority := properties.Actor{
		ObjectType: "Agent",
		Account: &properties.Account{
			Name:     "anonymous",
			HomePage: "http://cloud.scorm.com",
		},
	}

	stmt1 := statement.Statement{
		Actor:     actor,
		Verb:      verb,
		Context:   &context,
		Version:   utils.Ptr("1.0.0"),
		Authority: &authority,
		Object:    object,
	}

	s, err := utils.ToJson(stmt1, false)

	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(s), 0)

	s2, err := utils.ToJson(stmt1, false)

	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(s2), len(s))
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
