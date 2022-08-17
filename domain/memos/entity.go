package memos

import "time"

type Memo struct {
	ID             string
	DeviceToken    string
	YoutubeVideoID string
	Body           string
	Tags           []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
