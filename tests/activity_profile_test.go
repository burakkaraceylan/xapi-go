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

type ActivityProfileTestSuite struct {
	suite.Suite
	Activity statement.Activity
}

func (suite *ActivityProfileTestSuite) SetupSuite() {

	suite.Activity = *statement.NewActivityWithDefiniton(
		"http://github.com/bkaraceylan/xapi-go/Test/Unit/0",
		&statement.ActivityDefinition{
			Type:        utils.Ptr("http://github.com/bkaraceylan/xapi-go/activitytype/unit-test"),
			Name:        &statement.LanguageMap{"en-US": "Golang tests"},
			Description: &statement.LanguageMap{"en-US": "xapi-go golang client library unit tests"},
		},
	)
}

func (suite *ActivityProfileTestSuite) TestEmpty() {
	doc := documents.ActivityDocument{}

	assert.IsType(suite.T(), doc.Activity, statement.Activity{})

	assert.IsType(suite.T(), "", doc.ContentType)
	assert.IsType(suite.T(), []byte{}, doc.Content)
	assert.IsType(suite.T(), "", doc.Etag)
	assert.IsType(suite.T(), "", doc.ID)
	assert.IsType(suite.T(), time.Time{}, doc.Timestamp)
}

func (suite *ActivityProfileTestSuite) TestInitalized() {
	doc := documents.StateDocument{
		Activity: suite.Activity,
		Document: documents.Document{
			ID:          "test",
			ContentType: "test",
			Etag:        "test",
			Content:     []byte("test"),
			Timestamp:   time.Time{},
		},
	}

	assert.Equal(suite.T(), suite.Activity, doc.Activity)
	assert.Equal(suite.T(), "test", doc.ID)
	assert.Equal(suite.T(), "test", doc.ContentType)
	assert.Equal(suite.T(), "test", doc.Etag)
	assert.Equal(suite.T(), []byte("test"), doc.Content)
	assert.Equal(suite.T(), time.Time{}, doc.Timestamp)
}

func TestActivityProfileTestSuite(t *testing.T) {
	suite.Run(t, new(ActivityProfileTestSuite))
}
