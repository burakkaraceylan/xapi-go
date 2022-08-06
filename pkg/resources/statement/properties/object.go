package properties

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"

type InteractionComponet struct {
	ID          string               `json:"id" xapi:"required"`
	Description *special.LanguageMap `json:"description" xapi:"optional"`
}

// Object metadata
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#activity-definition
type Definition struct {
	Name                    *special.LanguageMap   `json:"name,omitempty" xApi:"recommended"`
	Description             *special.LanguageMap   `json:"description,omitempty" xApi:"recommended"`
	Type                    *string                `json:"type,omitempty" xApi:"recommended"`
	MoreInfo                *string                `json:"moreInfo,omitempty" xApi:"optional"`
	InteractionType         *string                `json:"interactionType,omitempty" xApi:"required"`
	CorrectResponsesPattern *[]string              `json:"correctResponsesPattern,omitempty" xApi:"optional"`
	Choices                 *[]InteractionComponet `json:"choices,omitempty" xApi:"optional"`
	Scale                   *[]InteractionComponet `json:"scale,omitempty" xApi:"optional"`
	Source                  *[]InteractionComponet `json:"source,omitempty" xApi:"optional"`
	Target                  *[]InteractionComponet `json:"target,omitempty" xApi:"optional"`
	Steps                   *[]InteractionComponet `json:"steps,omitempty" xApi:"optional"`
}

// The Object defines the thing that was acted on. The Object of a Statement can be an Activity, Agent/Group, SubStatement, or Statement
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#244-object
type Object struct {
	ID         string      `json:"id" xApi:"required"`
	ObjectType *string     `json:"objectType,omitempty" xApi:"optional"`
	Definition *Definition `json:"definition,omitempty" xApi:"optional"`
}
