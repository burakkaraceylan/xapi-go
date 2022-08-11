package activityprofile

import (
	"github.com/burakkaraceylan/xapi-go/pkg/resources"
	"github.com/burakkaraceylan/xapi-go/pkg/resources/statement/properties"
)

type ActivityDocument struct {
	resources.Document
	Activity properties.Object `json:"activityId"`
}
