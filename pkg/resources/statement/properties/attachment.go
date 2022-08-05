package properties

import "github.com/burakkaraceylan/xapi-go/pkg/resources/statement/special"

// In some cases an Attachment is logically an important part of a Learning Record.
// It could be an essay, a video, etc. Another example of such an Attachment is (the image of)
// a certificate that was granted as a result of an experience. It is useful to have a way to
// store these Attachments in and retrieve them from an LRS.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#2411-attachments

type Attachment struct {
	UsageType   string               `json:"usageType" xapi:"required"`
	Display     special.LanguageMap  `json:"display" xapi:"required"`
	Description *special.LanguageMap `json:"description,omitempty" xapi:"optional"`
	ContentType string               `json:"contentType" xapi:"required"`
	Length      int64                `json:"length" xapi:"required"`
	SHA2        string               `json:"sha1" xapi:"required"`
	FileUrl     *string              `json:"fileUrl,omitempty" xapi:"required"`
}
