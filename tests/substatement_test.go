package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SubstatementTestSuite struct {
	suite.Suite
	Agent    statement.Agent
	Activity statement.Activity
	Verb     statement.Verb
}

func (suite *SubstatementTestSuite) SetupSuite() {
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

func (suite *SubstatementTestSuite) TestMinimal() {
	ss := statement.NewSubStatement(suite.Agent, suite.Verb, suite.Activity)

	assert.IsType(suite.T(), &statement.SubStatement{}, ss)
	assert.Equal(suite.T(), suite.Agent, ss.Actor)
	assert.Equal(suite.T(), suite.Verb, ss.Verb)
	assert.Equal(suite.T(), suite.Activity, ss.Object)
}

func (suite *SubstatementTestSuite) TestWithOptionals() {
	attachment := statement.NewAttachment("test", statement.LanguageMap{"en-US": "test"}, "test", 10, "test")
	opt := statement.SubStatementOptions{
		Context: &statement.Context{
			Registration: utils.Ptr("test"),
		},
		Result: &statement.Result{
			Success: utils.Ptr(true),
		},
		Timestamp:   &time.Time{},
		Attachments: []statement.Attachment{*attachment},
	}

	ss := statement.NewSubStatement(suite.Agent, suite.Verb, suite.Activity, &opt)

	assert.IsType(suite.T(), &statement.SubStatement{}, ss)
	assert.Equal(suite.T(), suite.Agent, ss.Actor)
	assert.Equal(suite.T(), suite.Verb, ss.Verb)
	assert.Equal(suite.T(), suite.Activity, ss.Object)
	assert.Equal(suite.T(), time.Time{}, *ss.Timestamp)
	assert.Equal(suite.T(), "test", *ss.Context.Registration)
	assert.Equal(suite.T(), true, *ss.Result.Success)
	assert.Equal(suite.T(), 1, len(ss.Attachments))
	assert.Equal(suite.T(), *attachment, ss.Attachments[0])
}

func TestSubstatementTestSuite(t *testing.T) {
	suite.Run(t, new(SubstatementTestSuite))
}
