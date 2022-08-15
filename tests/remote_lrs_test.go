package tests

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/client"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/documents"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceTestSuite struct {
	suite.Suite
	lrs      *client.RemoteLRS
	Agent    statement.Agent
	Activity statement.Activity
	Verb     statement.Verb
}

func (suite *ResourceTestSuite) SetupSuite() {
	username := os.Getenv("XapiSandboxUsername")
	password := os.Getenv("XapiSandboxPassword")

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

	suite.Agent = *statement.NewAgentWithMbox("Burak Karaceylan", "mailto:bkaraceylan@gmail.com")

	suite.Activity = *statement.NewActivityWithDefiniton(
		"http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
		&statement.ActivityDefinition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &statement.LanguageMap{"en-US": "Golang tests"},
			Description: &statement.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		},
	)

	suite.Verb = statement.Verb{
		ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
		Display: statement.LanguageMap{"en-US": "Performed test"},
	}

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

	stmt := statement.NewStatement(suite.Agent, suite.Verb, suite.Activity)

	ids, resp, err := suite.lrs.SaveStatement(*stmt)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.NotEmpty(suite.T(), ids)
}

func (suite *ResourceTestSuite) TestSaveMultipleStatements() {
	stmt1 := statement.NewStatement(suite.Agent, suite.Verb, suite.Activity)
	stmt2 := statement.NewStatement(suite.Agent, suite.Verb, suite.Activity)

	stmts := []statement.Statement{*stmt1, *stmt2}

	ids, resp, err := suite.lrs.SaveStatements(stmts)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), len(ids), 2)
}

func (suite *ResourceTestSuite) TestPutStatement() {
	opt := statement.StatementOptions{
		ID: utils.Ptr(uuid.New().String()),
	}

	stmt := statement.NewStatement(suite.Agent, suite.Verb, suite.Activity, &opt)

	_, resp, err := suite.lrs.SaveStatement(*stmt)
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
	stmt := statement.NewStatement(suite.Agent, suite.Verb, suite.Activity)

	ids, resp, err := suite.lrs.SaveStatement(*stmt)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.NotEmpty(suite.T(), ids)

	retrieved, resp, err := suite.lrs.GetStatement(ids[0])
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), suite.Agent, *retrieved.Actor.(*statement.Agent))
	assert.Equal(suite.T(), suite.Activity, *retrieved.Object.(*statement.Activity))
	assert.Equal(suite.T(), suite.Verb, retrieved.Verb)
	assert.NotNil(suite.T(), retrieved.Stored)
	assert.NotEqual(suite.T(), time.Time{}, *retrieved.Stored)

	auth := statement.NewAgentWithAccount(
		"Unnamed Account",
		&statement.Account{
			HomePage: "http://cloud.scorm.com",
			Name:     "bKbNeGNaz-xbnTAHhR8",
		},
	)

	assert.Equal(suite.T(), *auth, *retrieved.Authority.(*statement.Agent))
}

func (suite *ResourceTestSuite) TestGetVoidedStatement() {
	// Test [POST]

	statement, resp, err := suite.lrs.GetVoidedStatement("6eae34d8-dd95-4609-af4c-214b85f359f7")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
	assert.Nil(suite.T(), statement)
}

func (suite *ResourceTestSuite) TestGetStateIds() {
	_, resp, err := suite.lrs.GetStateIds(suite.Activity, suite.Agent)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestGetState() {
	_, resp, err := suite.lrs.GetState(suite.Activity, suite.Agent, "test")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestSaveState() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.StateDocument{
		Agent:    suite.Agent,
		Activity: suite.Activity,
		Document: documents.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveState(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	rdoc, resp, err := suite.lrs.GetState(suite.Activity, suite.Agent, id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), id, rdoc.ID)
	assert.Equal(suite.T(), []byte("test"), rdoc.Content)
}

func (suite *ResourceTestSuite) TestDeleteState() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.StateDocument{
		Agent:    suite.Agent,
		Activity: suite.Activity,
		Document: documents.Document{
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

	doc := documents.StateDocument{
		Agent:    suite.Agent,
		Activity: suite.Activity,
		Document: documents.Document{
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

	ids, resp, err := suite.lrs.GetStateIds(suite.Activity, suite.Agent)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), 0, len(ids))
}

func (suite *ResourceTestSuite) TestGetActivityProfileIds() {
	_, resp, err := suite.lrs.GetActivityProfileIds(suite.Activity)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestGetActivityProfile() {
	_, resp, err := suite.lrs.GetActivityProfile(suite.Activity, "test")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestSaveActivtiyProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.ActivityDocument{
		Activity: suite.Activity,
		Document: documents.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveActivityProfile(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	rdoc, resp, err := suite.lrs.GetActivityProfile(suite.Activity, id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), id, rdoc.ID)
	assert.Equal(suite.T(), []byte("test"), rdoc.Content)
}

func (suite *ResourceTestSuite) TestDeleteActivityProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.ActivityDocument{
		Activity: suite.Activity,
		Document: documents.Document{
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

func (suite *ResourceTestSuite) TestGetAgentProfileIds() {
	_, resp, err := suite.lrs.GetAgentProfileIds(suite.Agent)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestGetAgentProfile() {
	_, resp, err := suite.lrs.GetAgentProfile(suite.Agent, "test")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.Response.StatusCode)
}

func (suite *ResourceTestSuite) TestSaveAgentProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.AgentDocument{
		Agent: suite.Agent,
		Document: documents.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveAgentProfile(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	rdoc, resp, err := suite.lrs.GetAgentProfile(suite.Agent, id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.Response.StatusCode)
	assert.Equal(suite.T(), id, rdoc.ID)
	assert.Equal(suite.T(), []byte("test"), rdoc.Content)
}

func (suite *ResourceTestSuite) TestDeleteAgentProfile() {
	rand.Seed(time.Now().UnixNano())

	id := RandomString(5)

	doc := documents.AgentDocument{
		Agent: suite.Agent,
		Document: documents.Document{
			ID:      id,
			Content: []byte("test"),
		},
	}

	_, resp, err := suite.lrs.SaveAgentProfile(&doc)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)

	resp, err = suite.lrs.DeleteAgentProfile(&doc)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 204, resp.Response.StatusCode)
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}
