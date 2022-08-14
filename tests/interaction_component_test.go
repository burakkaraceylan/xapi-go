package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InteractionComponentTestSuite struct {
	suite.Suite
}

func (suite *InteractionComponentTestSuite) TestMinimum() {
	ic := statement.NewInteractionComponent("test")

	assert.Equal(suite.T(), "test", ic.ID)
	assert.Nil(suite.T(), ic.Description)

}

func (suite *InteractionComponentTestSuite) TestOptionals() {

	ic := statement.NewInteractionComponent("test", &statement.LanguageMap{"en-US": "test"})

	assert.Equal(suite.T(), "test", ic.ID)
	assert.Equal(suite.T(), statement.LanguageMap{"en-US": "test"}, *ic.Description)
}

func TestInteractionComponentTestSuite(t *testing.T) {
	suite.Run(t, new(InteractionComponentTestSuite))
}
