package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/burakkaraceylan/xapi-go/pkg/client"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceTestSuite struct {
	suite.Suite
	lrs *client.RemoteLRS
}

func (suite *ResourceTestSuite) SetupSuite() {
	lrs, err := client.NewRemoteLRS(
		"https://cloud.scorm.com/ScormEngineInterface/TCAPI/public/",
		"1.0.0",
		"Basic VGVzdFVzZXI6cGFzc3dvcmQ=",
	)

	if err != nil {
		panic(err)
	}

	suite.lrs = lrs
}

func (suite *ResourceTestSuite) TestQueryParams() {
	actor := properties.Actor{
		ObjectType: "Test",
	}

	verb := properties.Verb{
		ID: "Test",
	}

	activity := properties.Object{
		ID: "Test",
	}

	params := client.QueryParams{
		StatementID:       utils.Ptr("Test"),
		VoidedStatementId: utils.Ptr("Test"),
		Agent:             &actor,
		Verb:              &verb,
		Activity:          &activity,
		Registeration:     utils.Ptr("Test"),
		RelatedActivities: utils.Ptr(true),
		RelatedAgents:     utils.Ptr(true),
		Since:             utils.Ptr(time.Time{}),
		Until:             utils.Ptr(time.Time{}),
		Format:            utils.Ptr("exact"),
		Attachments:       utils.Ptr(true),
		Ascending:         utils.Ptr(true),
	}

	q := params.Map()

	assert.Equal(suite.T(), "Test", q["statementId"])
	assert.Equal(suite.T(), "Test", q["voidedStatementId"])
	assert.Equal(suite.T(), "Test", q["verb"])
	assert.Equal(suite.T(), "Test", q["activity"])
	assert.Equal(suite.T(), "Test", q["registeration"])
	assert.Equal(suite.T(), "true", q["related_activities"])
	assert.Equal(suite.T(), "true", q["related_agents"])
	assert.Equal(suite.T(), "0001-01-01 00:00:00 +0000 UTC", q["since"])
	assert.Equal(suite.T(), "0001-01-01 00:00:00 +0000 UTC", q["until"])
	assert.Equal(suite.T(), "exact", q["format"])
	assert.Equal(suite.T(), "true", q["attachments"])
	assert.Equal(suite.T(), "true", q["ascending"])

}

func (suite *ResourceTestSuite) TestAboutResource() {
	// Test [GET]
	about, err := suite.lrs.About()

	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), about.Version)
}

// TODO: Test voided statement
func (suite *ResourceTestSuite) TestStatementResource() {
	// Test [POST]

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

	ids, resp, err := suite.lrs.SaveStatement(stmt1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.NotEmpty(suite.T(), ids)

	// Test [GET]

	stmt, resp, err := suite.lrs.GetStatement(ids[0])
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)

	assert.Equal(suite.T(), *stmt.ID, ids[0])
	assert.Equal(suite.T(), stmt.Actor.ObjectType, "Agent")
	assert.Equal(suite.T(), *stmt.Actor.Name, "Foo Bar")
	assert.Equal(suite.T(), *stmt.Actor.Mbox, "mailto:foo@bar.com")

	assert.Equal(suite.T(), stmt.Verb.ID, "http://adlnet.gov/expapi/verbs/initialized")
	assert.Equal(suite.T(), stmt.Verb.Display["en-US"], "initialized")

	assert.Equal(suite.T(), *stmt.Context.Registration, "e168d6a3-46b2-4233-82e7-66b73a179727")

	assert.Equal(suite.T(), len(*stmt.Context.ContextActivities.Category), 2)

	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Category)[0].ID, "http://id.tincanapi.com/recipe/tincan-prototypes/tetris/1")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[0].ObjectType, "Activity")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[0].Definition.Type, "http://id.tincanapi.com/activitytype/recipe")

	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Category)[1].ID, "http://id.tincanapi.com/activity/tincan-prototypes/tetris-template")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[1].ObjectType, "Activity")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[1].Definition.Type, "http://id.tincanapi.com/activitytype/source")

	assert.Equal(suite.T(), len(*stmt.Context.ContextActivities.Grouping), 2)

	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Grouping)[0].ID, "http://id.tincanapi.com/activity/tincan-prototypes")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[0].ObjectType, "Activity")

	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Grouping)[1].ID, "http://id.tincanapi.com/activity/tincan-prototypes/tetris")
	assert.Equal(suite.T(), *(*stmt.Context.ContextActivities.Category)[1].ObjectType, "Activity")

	assert.WithinDuration(suite.T(), time.Now(), *stmt.Timestamp, time.Second*10)
	assert.WithinDuration(suite.T(), time.Now(), *stmt.Stored, time.Second*10)

	assert.Equal(suite.T(), stmt.Authority.ObjectType, "Agent")
	assert.Equal(suite.T(), stmt.Authority.Account.Name, "anonymous")
	assert.Equal(suite.T(), stmt.Authority.Account.HomePage, "http://cloud.scorm.com")

	assert.Equal(suite.T(), *stmt.Version, "1.0.0")

	assert.Equal(suite.T(), stmt.Object.ID, "http://id.tincanapi.com/activity/tincan-prototypes/tetris")
	assert.Equal(suite.T(), (*stmt.Object.Definition.Name)["en-US"], "Js Tetris - Tin Can Prototype")
	assert.Equal(suite.T(), (*stmt.Object.Definition.Description)["en-US"], "A game of tetris.")
	assert.Equal(suite.T(), *stmt.Object.Definition.Type, "http://activitystrea.ms/schema/1.0/game")
	assert.Equal(suite.T(), *stmt.Object.ObjectType, "Activity")

	//Test Multiple [POST]
	stmts := []statement.Statement{stmt1, stmt1}
	ids, resp, err = suite.lrs.SaveStatements(stmts)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), len(ids), 2)

	//Test [PUT]
	stmt1.ID = utils.Ptr(uuid.New().String())
	_, resp, err = suite.lrs.SaveStatement(stmt1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	//Test Query [GET]
	params := client.QueryParams{
		Limit: utils.Ptr(int64(14)),
	}

	mresp, resp, err := suite.lrs.QueryStatements(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), len(mresp.Statements), 14)
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}
