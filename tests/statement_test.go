package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StatementTestSuite struct {
	suite.Suite
	Actor       statement.IActor
	Verb        statement.Verb
	Object      statement.IObject
	Result      statement.Result
	Context     statement.Context
	Timestamp   time.Time
	Stored      time.Time
	Authority   statement.IActor
	Version     string
	Attachments []statement.Attachment
}

func (suite *StatementTestSuite) SetupSuite() {
	suite.Actor = statement.NewAnonymousAgentWithMbox("mailto:bkaraceylan@gmail.com")

	suite.Object = statement.NewActivityWithDefiniton("http://github.com/bkaraceylan/xapi-go/Test/Unit/0",

		&statement.ActivityDefinition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &statement.LanguageMap{"en-US": "Golang tests"},
			Description: &statement.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		})

	suite.Verb = statement.Verb{
		ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
		Display: statement.LanguageMap{"en-US": "Performed test"},
	}

	suite.Result = statement.Result{
		Success: utils.Ptr(true),
	}

	suite.Context = statement.Context{
		Registration: utils.Ptr("test"),
	}

	suite.Timestamp = time.Time{}
	suite.Stored = time.Time{}
	suite.Authority = statement.NewAnonymousAgentWithMbox("mailto:bkaraceylan@gmail.com")
	suite.Version = "1.0.0"
	suite.Attachments = []statement.Attachment{}
}

func (suite *StatementTestSuite) TestNoOptions() {
	stmt := statement.NewStatement(suite.Actor, suite.Verb, suite.Object)

	str := ""
	strPtr := utils.Ptr(str)

	assert.IsType(suite.T(), strPtr, stmt.ID)
	assert.IsType(suite.T(), suite.Actor, stmt.Actor)
	assert.IsType(suite.T(), suite.Verb, stmt.Verb)
	assert.IsType(suite.T(), suite.Object, stmt.Object)
	assert.IsType(suite.T(), &statement.Result{}, stmt.Result)
	assert.IsType(suite.T(), &statement.Context{}, stmt.Context)
	assert.IsType(suite.T(), &time.Time{}, stmt.Timestamp)
	assert.IsType(suite.T(), &time.Time{}, stmt.Stored)
	assert.IsType(suite.T(), nil, stmt.Authority)
	assert.IsType(suite.T(), strPtr, stmt.Version)
	assert.IsType(suite.T(), []statement.Attachment{}, stmt.Attachments)
}

func (suite *StatementTestSuite) TestWithOptions() {

	id := utils.Ptr("test")

	opt := statement.StatementOptions{
		ID:          id,
		Result:      &suite.Result,
		Context:     &suite.Context,
		Timestamp:   &suite.Timestamp,
		Stored:      &suite.Stored,
		Authority:   suite.Authority,
		Version:     &suite.Version,
		Attachments: suite.Attachments,
	}

	stmt := statement.NewStatement(suite.Actor, suite.Verb, suite.Object, &opt)

	assert.Equal(suite.T(), id, stmt.ID)
	assert.Equal(suite.T(), suite.Actor, stmt.Actor)
	assert.Equal(suite.T(), suite.Object, stmt.Object)
	assert.Equal(suite.T(), suite.Verb, stmt.Verb)
	assert.Equal(suite.T(), &suite.Result, stmt.Result)
	assert.Equal(suite.T(), &suite.Context, stmt.Context)
	assert.Equal(suite.T(), &suite.Timestamp, stmt.Timestamp)
	assert.Equal(suite.T(), &suite.Stored, stmt.Stored)
	assert.Equal(suite.T(), suite.Authority, stmt.Authority)
	assert.Equal(suite.T(), &suite.Version, stmt.Version)
	assert.Equal(suite.T(), suite.Attachments, stmt.Attachments)
}

func TestStatementTestSuite(t *testing.T) {
	suite.Run(t, new(StatementTestSuite))
}
