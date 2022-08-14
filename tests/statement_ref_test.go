package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StatementRefTestSuite struct {
	suite.Suite
}

func (suite *StatementRefTestSuite) TestInitialize() {
	sr := statement.NewStatementRef("test")
	assert.IsType(suite.T(), &statement.StatementRef{}, sr)
	assert.Equal(suite.T(), "test", sr.ID)
	assert.Equal(suite.T(), "StatementRef", sr.ObjectType)
	assert.Equal(suite.T(), "StatementRef", sr.GetObjectType())
}
func TestStatementRefTestSuite(t *testing.T) {
	suite.Run(t, new(StatementRefTestSuite))
}
