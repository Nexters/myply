package memos

import "time"

type Memo struct {
	Id             string
	DeviceToken    string
	YoutubeVideoId string
	Body           string
	Tags           []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
