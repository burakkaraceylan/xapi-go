package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	lrs *RemoteLRS
}

func (suite *TestSuite) SetupSuite() {
	lrs, err := NewRemoteLRS(
		"https://cloud.scorm.com/ScormEngineInterface/TCAPI/public/",
		"1.0.0",
		"Basic VGVzdFVzZXI6cGFzc3dvcmQ=",
	)

	if err != nil {
		panic(err)
	}

	suite.lrs = lrs
}

func (suite *TestSuite) TestAbout() {
	about, err := suite.lrs.About()

	assert.Nil(suite.T(), err)
	assert.NotEmpty(suite.T(), about.Version)
}

// ? Statement used in the test may change
func (suite *TestSuite) TestGetStatement() {
	stmt, err := suite.lrs.GetStatement("707a44e8-0163-42a2-a86a-bd50f569d8d5")
	fmt.Printf(*stmt.Context.Registration)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), stmt.ID, "707a44e8-0163-42a2-a86a-bd50f569d8d5")
	assert.Equal(suite.T(), stmt.Actor.ObjectType, "Agent")
	assert.Equal(suite.T(), *stmt.Actor.Name, "Test User")
	assert.Equal(suite.T(), *stmt.Actor.Mbox, "mailto:test@beta.projecttincan.com")

	assert.Equal(suite.T(), stmt.Verb.ID, "http://adlnet.gov/expapi/verbs/answered")
	assert.Equal(suite.T(), stmt.Verb.Display["en-US"], "answered")

	assert.Equal(suite.T(), *stmt.Result.Success, true)
	assert.Equal(suite.T(), *stmt.Result.Response, "")

	assert.Equal(suite.T(), *stmt.Context.Registration, "e168d6a3-46b2-4233-82e7-66b73a179727")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Parent[0].ID, "http://id.tincanapi.com/activity/tincan-prototypes/golf-example/GolfAssessment")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Parent[0].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Category[0].ID, "http://id.tincanapi.com/recipe/tincan-prototypes/golf/1")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Category[0].Definition.Type, "http://id.tincanapi.com/activitytype/recipe")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Category[0].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Category[1].ID, "http://id.tincanapi.com/activity/tincan-prototypes/elearning")
	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Category[1].Definition.Name)["en-US"], "E-learning course")
	assert.Equal(suite.T(), (*stmt.Context.ContextActivities.Category[1].Definition.Description)["en-US"], "An e-learning course built using the golf prototype framework.")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Category[1].Definition.Type, "http://id.tincanapi.com/activitytype/source")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Category[1].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Grouping[0].ID, "http://id.tincanapi.com/activity/tincan-prototypes/golf-example")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Grouping[0].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Grouping[1].ID, "http://id.tincanapi.com/activity/tincan-prototypes")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Grouping[1].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Context.ContextActivities.Grouping[2].ID, "http://id.tincanapi.com/activity/tincan-prototypes/golf-example/GolfAssessment")
	assert.Equal(suite.T(), *stmt.Context.ContextActivities.Grouping[2].ObjectType, "Activity")
	assert.Equal(suite.T(), stmt.Timestamp.Format(time.RFC3339Nano), "2022-08-05T03:52:10.682Z")
	assert.Equal(suite.T(), stmt.Stored.Format(time.RFC3339Nano), "2022-08-05T03:52:11.041Z")
	assert.Equal(suite.T(), stmt.Authority.ObjectType, "Agent")
	assert.Equal(suite.T(), stmt.Authority.Account.Name, "anonymous")
	assert.Equal(suite.T(), stmt.Authority.Account.HomePage, "http://cloud.scorm.com")
	assert.Equal(suite.T(), *stmt.Version, "1.0.0")
	assert.Equal(suite.T(), stmt.Object.ID, "http://id.tincanapi.com/activity/tincan-prototypes/golf-example/GolfAssessment/interactions.handicap_3")
	assert.Equal(suite.T(), (*stmt.Object.Definition.Description)["en-US"], "A 'scratch golfer' has a handicap of ___")
	assert.Equal(suite.T(), *stmt.Object.Definition.Type, "http://adlnet.gov/expapi/activities/cmi.interaction")
	assert.Equal(suite.T(), (*stmt.Object.Definition.CorrectResponsesPattern)[0], "0[:]0")
	assert.Equal(suite.T(), *stmt.Object.ObjectType, "Activity")

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
