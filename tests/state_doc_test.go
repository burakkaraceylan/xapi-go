package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/documents"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StateDocTestSuite struct {
	suite.Suite
	Agent    statement.Agent
	Activity statement.Activity
}

func (suite *StateDocTestSuite) SetupSuite() {
	suite.Agent = statement.Agent{
		ObjectType: "Agent",
		Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
	}

	suite.Activity = *statement.NewActivityWithDefiniton("http://github.com/bkaraceylan/xapi-go/Test/Unit/0", &statement.ActivityDefinition{
		Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
		Name:        &statement.LanguageMap{"en-US": "Golang tests"},
		Description: &statement.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
	})

}

func (suite *StateDocTestSuite) TestEmpty() {
	doc := documents.StateDocument{}

	assert.IsType(suite.T(), doc.Agent, statement.Agent{})
	assert.IsType(suite.T(), doc.Activity, statement.Activity{})
	assert.IsType(suite.T(), doc.Activity, statement.Activity{})

	str := ""
	strPtry := utils.Ptr(str)

	assert.IsType(suite.T(), strPtry, doc.Registration)
	assert.IsType(suite.T(), "", doc.ContentType)
	assert.IsType(suite.T(), []byte{}, doc.Content)
	assert.IsType(suite.T(), "", doc.Etag)
	assert.IsType(suite.T(), "", doc.ID)
	assert.IsType(suite.T(), time.Time{}, doc.Timestamp)
}

func (suite *StateDocTestSuite) TestInitalized() {
	doc := documents.StateDocument{
		Agent:    suite.Agent,
		Activity: suite.Activity,
		Document: documents.Document{
			ID:          "test",
			ContentType: "test",
			Etag:        "test",
			Content:     []byte("test"),
			Timestamp:   time.Time{},
		},
	}

	assert.Equal(suite.T(), suite.Agent, doc.Agent)
	assert.Equal(suite.T(), suite.Activity, doc.Activity)
	assert.Equal(suite.T(), "test", doc.ID)
	assert.Equal(suite.T(), "test", doc.ContentType)
	assert.Equal(suite.T(), "test", doc.Etag)
	assert.Equal(suite.T(), []byte("test"), doc.Content)
	assert.Equal(suite.T(), time.Time{}, doc.Timestamp)
}

func TestStateDocTestSuite(t *testing.T) {
	suite.Run(t, new(StateDocTestSuite))
}
