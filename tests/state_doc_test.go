package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/state"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StateDocTestSuite struct {
	suite.Suite
	Agent    properties.Actor
	Activity properties.Object
}

func (suite *StateDocTestSuite) SetupSuite() {
	suite.Agent = properties.Actor{
		ObjectType: "Agent",
		Mbox:       utils.Ptr("mailto:bkaraceylan@gmail.com"),
	}

	suite.Activity = properties.Object{
		ObjectType: utils.Ptr("Activity"),
		ID:         "http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
		Definition: &properties.Definition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &special.LanguageMap{"en-US": "Golang tests"},
			Description: &special.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		},
	}
}

func (suite *StateDocTestSuite) TestEmpty() {
	doc := state.StateDocument{}

	assert.IsType(suite.T(), doc.Agent, properties.Actor{})
	assert.IsType(suite.T(), doc.Activity, properties.Object{})
	assert.IsType(suite.T(), doc.Activity, properties.Object{})

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
	doc := state.StateDocument{
		Agent:    suite.Agent,
		Activity: suite.Activity,
		Document: resources.Document{
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
