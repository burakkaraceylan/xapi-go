package documents

import "time"

// DOcument is the base trype for document resources
type Document struct {
	ID          string    `json:"-"`
	ContentType string    `json:"-"`
	Content     []byte    `json:"-"`
	Etag        string    `json:"-"`
	Timestamp   time.Time `json:"-"`
}
