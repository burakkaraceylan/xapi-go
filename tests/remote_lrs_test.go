package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/client"
	"github.com/burakkaraceylan/xapi-go/pkg/resources"
	activityprofile "github.com/burakkaraceylan/xapi-go/pkg/resources/activity_profile"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/state"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceTestSuite struct {
	suite.Suite
	lrs    *client.RemoteLRS
	Agent  properties.Actor
	Object properties.Object
	Verb   properties.Verb
}

func (suite *ResourceTestSuite) SetupSuite() {
	//username := os.Getenv("XapiSandboxUsername")
	//password := os.Getenv("XapiSandboxPassword")

	username := "bKbNeGNaz-xbnTAHhR8"
	password := "mA1QI0wZuaBs2HBzCQI"

	if len(username) == 0 {
		panic("Sandbox username not present in environment variables")
	}

	if len(password) == 0 {
		panic("Sandbox password not present in environment variables")
	}

	lrs, err := client.NewRemoteLRS(
		"https://cloud.scorm.com/lrs/OSQO3KVP5L/sandbox/",
		"1.0.0",
		username,
		password,
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

	params := client.StatementQueryParams{
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
func (suite *ResourceTestSuite) TestSaveStatement() {
	// Test [POST]

	stmt := statement.Statement{
		Actor: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Object: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Verb: properties.Verb{
			ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
			Display: special.LanguageMap{"en-US": "Performed test"},
		},
	}

	ids, resp, err := suite.lrs.SaveStatement(stmt)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.NotEmpty(suite.T(), ids)
}

func (suite *ResourceTestSuite) TestSaveMultipleStatements() {
	stmt1 := statement.Statement{
		Actor: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Object: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Verb: properties.Verb{
			ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
			Display: special.LanguageMap{"en-US": "Performed test"},
		},
	}

	stmt2 := statement.Statement{
		Actor: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Object: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Verb: properties.Verb{
			ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
			Display: special.LanguageMap{"en-US": "Performed test"},
		},
	}

	stmts := []statement.Statement{stmt1, stmt2}

	ids, resp, err := suite.lrs.SaveStatements(stmts)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), len(ids), 2)
}

func (suite *ResourceTestSuite) TestPutStatement() {
	stmt := statement.Statement{
		ID: utils.Ptr(uuid.New().String()),
		Actor: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Object: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Verb: properties.Verb{
			ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
			Display: special.LanguageMap{"en-US": "Performed test"},
		},
	}

	_, resp, err := suite.lrs.SaveStatement(stmt)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestStatementQueryParams() {
	params := client.StatementQueryParams{
		Limit: utils.Ptr(int64(14)),
	}

	mresp, resp, err := suite.lrs.QueryStatements(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.LessOrEqual(suite.T(), len(mresp.Statements), 14)
}

func (suite *ResourceTestSuite) TestGetStatement() {
	// Test [POST]

	actor := properties.Actor{
		ObjectType: "Agent",
		Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
	}

	object := properties.Object{
		ObjectType: utils.Ptr("Activity"),
		ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
		Definition: &properties.Definition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &special.LanguageMap{"en-US": "Golang tests"},
			Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		},
	}

	verb := properties.Verb{
		ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
		Display: special.LanguageMap{"en-US": "Performed test"},
	}

	stmt := statement.Statement{
		Actor:  actor,
		Object: object,
		Verb:   verb,
	}

	ids, resp, err := suite.lrs.SaveStatement(stmt)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.NotEmpty(suite.T(), ids)

	retrieved, resp, err := suite.lrs.GetStatement(ids[0])
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), actor, retrieved.Actor)
	assert.Equal(suite.T(), object, retrieved.Object)
	assert.Equal(suite.T(), verb, retrieved.Verb)
	assert.NotNil(suite.T(), retrieved.Stored)
	assert.NotEqual(suite.T(), time.Time{}, *retrieved.Stored)

	auth := properties.Actor{
		ObjectType: "Agent",
		Name:       utils.Ptr("Unnamed Account"),
		Account: &properties.Account{
			HomePage: "http://cloud.scorm.com",
			Name:     "bKbNeGNaz-xbnTAHhR8",
		},
	}

	assert.Equal(suite.T(), auth, *retrieved.Authority)
}

func (suite *ResourceTestSuite) TestGetStateIds() {
	params := client.GetStateIdsQueryParams{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
	}

	_, resp, err := suite.lrs.GetStateIds(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestGetState() {
	params := client.GetStateQueryParams{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		StateID: "test",
	}

	_, resp, err := suite.lrs.GetState(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestSaveState() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := state.StateDocument{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Document: resources.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveState(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	params := client.GetStateQueryParams{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		StateID: id,
	}

	rdoc, resp, err := suite.lrs.GetState(&params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), id, rdoc.ID)
	assert.Equal(suite.T(), []byte("test"), rdoc.Content)
}

func (suite *ResourceTestSuite) TestDeleteState() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := state.StateDocument{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Document: resources.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveState(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	resp, err = suite.lrs.DeleteState(&doc)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestClearStates() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := state.StateDocument{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Document: resources.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveState(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	doc.ID = RandomString(5)

	_, resp, err = suite.lrs.SaveState(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	doc.ID = ""

	resp, err = suite.lrs.DeleteState(&doc)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	query := client.GetStateIdsQueryParams{
		Agent: properties.Actor{
			ObjectType: "Agent",
			Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
		},
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
	}

	ids, resp, err := suite.lrs.GetStateIds(&query)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), 0, len(ids))
}

func (suite *ResourceTestSuite) TestGetActivityProfileIds() {
	params := client.GetActivityProfilesParams{
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
	}

	_, resp, err := suite.lrs.GetActivityProfileIds(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestGetActivityProfile() {
	params := client.GetActivityProfileParams{
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		ProfileID: "test",
	}

	_, resp, err := suite.lrs.GetActivityProfile(&params)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestSaveActivtiyProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := activityprofile.ActivityDocument{
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Document: resources.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveActivityProfile(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	params := client.GetActivityProfileParams{
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		ProfileID: id,
	}

	rdoc, resp, err := suite.lrs.GetActivityProfile(&params)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), id, rdoc.ID)
	assert.Equal(suite.T(), []byte("test"), rdoc.Content)
}

func (suite *ResourceTestSuite) TestDeleteActivityProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := activityprofile.ActivityDocument{
		Activity: properties.Object{
			ObjectType: utils.Ptr("Activity"),
			ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
			Definition: &properties.Definition{
				Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
				Name:        &special.LanguageMap{"en-US": "Golang tests"},
				Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
			},
		},
		Document: resources.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveActivityProfile(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	resp, err = suite.lrs.DeleteActivityProfile(&doc)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}
