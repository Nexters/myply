package memos

import (
	"time"
)

type Memo struct {
	ID             string
	DeviceToken    string
	YoutubeVideoID string
	Body           string
	Tags           []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Memos []Memo

func (ms *Memos) YoutubeVideoIDs() []string {
	var youtubeIDs []string
	for _, m := range *ms {
		youtubeIDs = append(youtubeIDs, m.YoutubeVideoID)
	}
	return youtubeIDs
}
