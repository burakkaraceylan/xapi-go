package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StatementTestSuite struct {
	suite.Suite
	Actor       properties.Actor
	Verb        properties.Verb
	Object      properties.Object
	Result      properties.Result
	Context     properties.Context
	Timestamp   time.Time
	Stored      time.Time
	Authority   properties.Actor
	Version     string
	Attachments []properties.Attachment
}

func (suite *StatementTestSuite) SetupSuite() {
	suite.Actor = properties.Actor{
		ObjectType: "Agent",
		Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
	}

	suite.Object = properties.Object{
		ObjectType: utils.Ptr("Activity"),
		ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
		Definition: &properties.Definition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &special.LanguageMap{"en-US": "Golang tests"},
			Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		},
	}

	suite.Verb = properties.Verb{
		ID:      "http://github.com/bkaraceylan/xapi-go/Test/performed",
		Display: special.LanguageMap{"en-US": "Performed test"},
	}

	suite.Result = properties.Result{
		Success: utils.Ptr(true),
	}

	suite.Context = properties.Context{
		Registration: utils.Ptr("test"),
	}

	suite.Timestamp = time.Time{}
	suite.Stored = time.Time{}
	suite.Authority = properties.Actor{
		ObjectType: "Agent",
		Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
	}
	suite.Version = "1.0.0"
	suite.Attachments = []properties.Attachment{}
}

func (suite *StatementTestSuite) TestEmpty() {
	stmt := statement.Statement{}

	str := ""
	strPtr := utils.Ptr(str)

	assert.IsType(suite.T(), strPtr, stmt.ID)
	assert.IsType(suite.T(), properties.Actor{}, stmt.Actor)
	assert.IsType(suite.T(), properties.Verb{}, stmt.Verb)
	assert.IsType(suite.T(), properties.Object{}, stmt.Object)
	assert.IsType(suite.T(), &properties.Result{}, stmt.Result)
	assert.IsType(suite.T(), &properties.Context{}, stmt.Context)
	assert.IsType(suite.T(), &time.Time{}, stmt.Timestamp)
	assert.IsType(suite.T(), &time.Time{}, stmt.Stored)
	assert.IsType(suite.T(), &properties.Actor{}, stmt.Authority)
	assert.IsType(suite.T(), strPtr, stmt.Version)
	assert.IsType(suite.T(), &[]properties.Attachment{}, stmt.Attachments)
}

func (suite *StatementTestSuite) TestInitialize() {

	id := utils.Ptr("test")

	stmt := statement.Statement{
		ID:          id,
		Actor:       suite.Actor,
		Object:      suite.Object,
		Verb:        suite.Verb,
		Result:      &suite.Result,
		Context:     &suite.Context,
		Timestamp:   &suite.Timestamp,
		Stored:      &suite.Stored,
		Authority:   &suite.Authority,
		Version:     &suite.Version,
		Attachments: &suite.Attachments,
	}

	assert.Equal(suite.T(), id, stmt.ID)
	assert.Equal(suite.T(), suite.Actor, stmt.Actor)
	assert.Equal(suite.T(), suite.Object, stmt.Object)
	assert.Equal(suite.T(), suite.Verb, stmt.Verb)
	assert.Equal(suite.T(), &suite.Result, stmt.Result)
	assert.Equal(suite.T(), &suite.Context, stmt.Context)
	assert.Equal(suite.T(), &suite.Timestamp, stmt.Timestamp)
	assert.Equal(suite.T(), &suite.Stored, stmt.Stored)
	assert.Equal(suite.T(), &suite.Authority, stmt.Authority)
	assert.Equal(suite.T(), &suite.Version, stmt.Version)
	assert.Equal(suite.T(), &suite.Attachments, stmt.Attachments)
}

func TestStatementTestSuite(t *testing.T) {
	suite.Run(t, new(StatementTestSuite))
}
