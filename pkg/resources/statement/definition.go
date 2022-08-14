package statement

import "errors"

// Object metadata
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#activity-definition
type ActivityDefinition struct {
	Name                    *LanguageMap           `json:"name,omitempty" xApi:"recommended"`
	Description             *LanguageMap           `json:"description,omitempty" xApi:"recommended"`
	Type                    *string                `json:"type,omitempty" xApi:"recommended"`
	MoreInfo                *string                `json:"moreInfo,omitempty" xApi:"optional"`
	InteractionType         *string                `json:"interactionType,omitempty" xApi:"required"`
	CorrectResponsesPattern []string               `json:"correctResponsesPattern,omitempty" xApi:"optional"`
	Choices                 []InteractionComponent `json:"choices,omitempty" xApi:"optional"`
	Scale                   []InteractionComponent `json:"scale,omitempty" xApi:"optional"`
	Source                  []InteractionComponent `json:"source,omitempty" xApi:"optional"`
	Target                  []InteractionComponent `json:"target,omitempty" xApi:"optional"`
	Steps                   []InteractionComponent `json:"steps,omitempty" xApi:"optional"`
}

func (ad *ActivityDefinition) AddChoice(component InteractionComponent) error {
	if *ad.InteractionType != "choice" && *ad.InteractionType != "sequencing" {
		return errors.New("interaction type must be choice or sequencing")
	}

	ad.Choices = append(ad.Choices, component)
	return nil
}

func (ad *ActivityDefinition) AddScale(component InteractionComponent) error {
	if *ad.InteractionType != "likert" {
		return errors.New("interaction type must be likert")
	}

	ad.Scale = append(ad.Scale, component)
	return nil
}

func (ad *ActivityDefinition) AddMatching(component InteractionComponent, isTarget bool) error {
	if *ad.InteractionType != "matching" {
		return errors.New("interaction type must be matching")
	}

	if isTarget {
		ad.Target = append(ad.Target, component)
		return nil
	}

	ad.Source = append(ad.Source, component)
	return nil
}

func (ad *ActivityDefinition) AddSteps(component InteractionComponent) error {
	if *ad.InteractionType != "performance" {
		return errors.New("interaction type must be performance")
	}

	ad.Steps = append(ad.Steps, component)
	return nil
}
