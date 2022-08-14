package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/about"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AboutTestResource struct {
	suite.Suite
}

func (suite *AboutTestResource) TestEmpty() {
	a := about.About{}

	assert.IsType(suite.T(), []string{}, a.Version)
	assert.IsType(suite.T(), &statement.Extensions{}, a.Extensions)
}

func (suite *AboutTestResource) TestInitalized() {
	a := about.About{
		Version: []string{"0.0.9", "1.0.0"},
		Extensions: &statement.Extensions{
			"test-extension-1": "https://github.com/bkaraceylan/xapi-go/ext/1",
			"test-extension-2": "https://github.com/bkaraceylan/xapi-go/ext/2",
		},
	}

	assert.Equal(suite.T(), "0.0.9", a.Version[0])
	assert.Equal(suite.T(), "1.0.0", a.Version[1])

	assert.Contains(suite.T(), "test-extension-1", a.Extensions)
	assert.Contains(suite.T(), "test-extension-2", a.Extensions)
}

func TestAboutTestResource(t *testing.T) {
	suite.Run(t, new(StateDocTestSuite))
}
