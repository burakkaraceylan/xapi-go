package tests

import (
	"testing"

	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DefinitionTestSuite struct {
	suite.Suite
}

func (suite *DefinitionTestSuite) TestEmpty() {
	d := statement.ActivityDefinition{}

	assert.Nil(suite.T(), d.Name)
	assert.Nil(suite.T(), d.Choices)
	assert.Nil(suite.T(), d.CorrectResponsesPattern)
	assert.Nil(suite.T(), d.Description)
	assert.Nil(suite.T(), d.InteractionType)
	assert.Nil(suite.T(), d.MoreInfo)
	assert.Nil(suite.T(), d.Scale)
	assert.Nil(suite.T(), d.Source)
	assert.Nil(suite.T(), d.Steps)
	assert.Nil(suite.T(), d.Target)
	assert.Nil(suite.T(), d.Type)
}

func (suite *DefinitionTestSuite) TestInitialized() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "choice"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	_ = d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	_ = d.AddChoice(*statement.NewInteractionComponent("choice2", &statement.LanguageMap{"en-US": "test"}))

	assert.Equal(suite.T(), name, *d.Name)
	assert.Equal(suite.T(), atype, *d.Type)
	assert.Equal(suite.T(), description, *d.Description)
	assert.Equal(suite.T(), minfo, *d.MoreInfo)
	assert.Equal(suite.T(), itype, *d.InteractionType)
	assert.Equal(suite.T(), crpattern, d.CorrectResponsesPattern)

	assert.Equal(suite.T(), 2, len(d.Choices))
	assert.Nil(suite.T(), d.Scale)
	assert.Nil(suite.T(), d.Source)
	assert.Nil(suite.T(), d.Steps)
	assert.Nil(suite.T(), d.Target)
}

func (suite *DefinitionTestSuite) TestInteractionTypeChoice() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "choice"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	err := d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	assert.Nil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), false)
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), true)
	assert.NotNil(suite.T(), err)

	err = d.AddScale(*statement.NewInteractionComponent("scale", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddSteps(*statement.NewInteractionComponent("steps", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)
}

func (suite *DefinitionTestSuite) TestInteractionTypeSequencing() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "sequencing"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	err := d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	assert.Nil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), false)
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), true)
	assert.NotNil(suite.T(), err)

	err = d.AddScale(*statement.NewInteractionComponent("scale", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddSteps(*statement.NewInteractionComponent("steps", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)
}

func (suite *DefinitionTestSuite) TestInteractionTypeLikert() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "likert"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	err := d.AddScale(*statement.NewInteractionComponent("scale", &statement.LanguageMap{"en-US": "test"}))
	assert.Nil(suite.T(), err)

	err = d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), false)
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), true)
	assert.NotNil(suite.T(), err)

	err = d.AddSteps(*statement.NewInteractionComponent("steps", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)
}

func (suite *DefinitionTestSuite) TestInteractionTypeMatching() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "matching"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	err := d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), false)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(d.Source))
	assert.Equal(suite.T(), 0, len(d.Target))

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), true)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(d.Source))
	assert.Equal(suite.T(), 1, len(d.Target))

	err = d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddScale(*statement.NewInteractionComponent("scale", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddSteps(*statement.NewInteractionComponent("steps", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)
}

func (suite *DefinitionTestSuite) TestInteractionTypePerformance() {
	name := statement.LanguageMap{"en-US": "test"}
	description := statement.LanguageMap{"en-US": "test"}
	atype := "test"
	minfo := "test"
	itype := "performance"
	crpattern := []string{"test", "test"}

	d := statement.ActivityDefinition{
		Name:                    &name,
		Description:             &description,
		Type:                    &atype,
		MoreInfo:                &minfo,
		InteractionType:         &itype,
		CorrectResponsesPattern: crpattern,
	}

	err := d.AddSteps(*statement.NewInteractionComponent("steps", &statement.LanguageMap{"en-US": "test"}))
	assert.Nil(suite.T(), err)

	err = d.AddChoice(*statement.NewInteractionComponent("choice1", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), false)
	assert.NotNil(suite.T(), err)

	err = d.AddMatching(*statement.NewInteractionComponent("match", &statement.LanguageMap{"en-US": "test"}), true)
	assert.NotNil(suite.T(), err)

	err = d.AddScale(*statement.NewInteractionComponent("scale", &statement.LanguageMap{"en-US": "test"}))
	assert.NotNil(suite.T(), err)
}

func TestDefinitionTestSuite(t *testing.T) {
	suite.Run(t, new(DefinitionTestSuite))
}
