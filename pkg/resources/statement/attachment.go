package statement

// In some cases an Attachment is logically an important part of a Learning Record.
// It could be an essay, a video, etc. Another example of such an Attachment is (the image of)
// a certificate that was granted as a result of an experience. It is useful to have a way to
// store these Attachments in and retrieve them from an LRS.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#2411-attachments

type Attachment struct {
	UsageType   string      `json:"usageType" xapi:"required"`
	Display     LanguageMap `json:"display" xapi:"required"`
	ContentType string      `json:"contentType" xapi:"required"`
	Length      int64       `json:"length" xapi:"required"`
	SHA2        string      `json:"sha1" xapi:"required"`
	AttachmentOptions
}

type AttachmentOptions struct {
	Description *LanguageMap `json:"description,omitempty" xapi:"optional"`
	FileUrl     *string      `json:"fileUrl,omitempty" xapi:"optional"`
}

func NewAttachment(usageType string, display LanguageMap, contentType string, length int64, sha2 string, params ...*AttachmentOptions) *Attachment {
	var opt *AttachmentOptions

	if len(params) == 1 {
		opt = params[0]
	}

	attachment := Attachment{
		UsageType:   usageType,
		Display:     display,
		ContentType: contentType,
		Length:      length,
		SHA2:        sha2,
	}

	if opt != nil {
		if opt.Description != nil {
			attachment.Description = opt.Description
		}

		if opt.Description != nil {
			attachment.FileUrl = opt.FileUrl
		}
	}

	return &attachment
}
