package resources

import "time"

type Document struct {
	ID          string    `json:"-"`
	ContentType string    `json:"-"`
	Content     []byte    `json:"-"`
	Etag        string    `json:"-"`
	Timestamp   time.Time `json:"-"`
}
