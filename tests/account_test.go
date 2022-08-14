package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite
}

func (suite *AccountTestSuite) TestInitialize() {
	a := statement.NewAccount("http://github.com", "bkaraceylan")

	assert.Equal(suite.T(), "http://github.com", a.HomePage)
	assert.Equal(suite.T(), "bkaraceylan", a.Name)

}
func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
