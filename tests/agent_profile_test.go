package tests

import (
	"testing"
	"time"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/documents"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AgentProfileTestSuite struct {
	suite.Suite
	Agent statement.Agent
}

func (suite *AgentProfileTestSuite) SetupSuite() {

	suite.Agent = *statement.NewAnonymousAgentWithMbox("mailto:bkaraceylan@gmail.com")
}

func (suite *AgentProfileTestSuite) TestEmpty() {
	doc := documents.AgentDocument{}

	assert.IsType(suite.T(), doc.Agent, statement.Agent{})

	assert.IsType(suite.T(), "", doc.ContentType)
	assert.IsType(suite.T(), []byte{}, doc.Content)
	assert.IsType(suite.T(), "", doc.Etag)
	assert.IsType(suite.T(), "", doc.ID)
	assert.IsType(suite.T(), time.Time{}, doc.Timestamp)
}

func (suite *AgentProfileTestSuite) TestInitalized() {
	doc := documents.StateDocument{
		Agent: suite.Agent,
		Document: documents.Document{
			ID:          "test",
			ContentType: "test",
			Etag:        "test",
			Content:     []byte("test"),
			Timestamp:   time.Time{},
		},
	}

	assert.Equal(suite.T(), suite.Agent, doc.Agent)
	assert.Equal(suite.T(), "test", doc.ID)
	assert.Equal(suite.T(), "test", doc.ContentType)
	assert.Equal(suite.T(), "test", doc.Etag)
	assert.Equal(suite.T(), []byte("test"), doc.Content)
	assert.Equal(suite.T(), time.Time{}, doc.Timestamp)
}

func TestAgentProfileTestSuite(t *testing.T) {
	suite.Run(t, new(AgentProfileTestSuite))
}
