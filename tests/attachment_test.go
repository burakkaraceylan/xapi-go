package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/burakkaraceylan/xapi-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AttachmentTestSuite struct {
	suite.Suite
}

func (suite *AttachmentTestSuite) TestMinimum() {
	a := statement.NewAttachment("test", statement.LanguageMap{"en-US": "test"}, "test", 10, "test")

	assert.Equal(suite.T(), "test", a.UsageType)
	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, a.Display)
	assert.Equal(suite.T(), "test", a.ContentType)
	assert.Equal(suite.T(), int64(10), a.Length)
	assert.Equal(suite.T(), "test", a.SHA2)

	assert.Nil(suite.T(), a.Description)
	assert.Nil(suite.T(), a.FileUrl)
}

func (suite *AttachmentTestSuite) TestWithOptionals() {
	opt := statement.AttachmentOptions{
		Description: &statement.LanguageMap{"en-US": "test"},
		FileUrl:     utils.Ptr("test"),
	}
	a := statement.NewAttachment("test", statement.LanguageMap{"en-US": "test"}, "test", 10, "test", &opt)

	assert.Equal(suite.T(), "test", a.UsageType)
	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, a.Display)
	assert.Equal(suite.T(), "test", a.ContentType)
	assert.Equal(suite.T(), int64(10), a.Length)
	assert.Equal(suite.T(), "test", a.SHA2)

	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, *a.Description)
	assert.Equal(suite.T(), "test", *a.FileUrl)
}

func (suite *AttachmentTestSuite) TestInvalidOptionals() {
	a := statement.NewAttachment("test", statement.LanguageMap{"en-US": "test"}, "test", 10, "test", nil)

	assert.Equal(suite.T(), "test", a.UsageType)
	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, a.Display)
	assert.Equal(suite.T(), "test", a.ContentType)
	assert.Equal(suite.T(), int64(10), a.Length)
	assert.Equal(suite.T(), "test", a.SHA2)

	assert.Nil(suite.T(), a.Description)
	assert.Nil(suite.T(), a.FileUrl)
}

func TestAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(AttachmentTestSuite))
}
