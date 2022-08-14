package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type VerbTestSuite struct {
	suite.Suite
}

func (suite *VerbTestSuite) TestInitialized() {
	v := statement.NewVerb("test", statement.LanguageMap{"en-US": "test"})

	assert.IsType(suite.T(), "", v.ID)
	assert.IsType(suite.T(), statement.LanguageMap{}, v.Display)

	assert.Equal(suite.T(), "test", v.ID)
	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, v.Display)
}

func TestVerbTestSuite(t *testing.T) {
	suite.Run(t, new(VerbTestSuite))
}
