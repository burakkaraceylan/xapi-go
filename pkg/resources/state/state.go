package state

import (
	"github.com/burakkaraceylan/xapi-go/pkg/resources"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

type StateDocument struct {
	resources.Document
	Activity     properties.Object `json:"activityId"`
	Agent        properties.Actor  `json:"agent"`
	Registration *string           `json:"registration"`
}
