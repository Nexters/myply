package memos

import (
	"time"
)

type Memo struct {
	Id             string
	DeviceToken    string
	YoutubeVideoId string
	Body           string
	TagIds         []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Repository interface {
	GetMemo(id string) (*Memo, error)
	SaveMemo(videoId string, body string, deviceToken string) (string, error)
	GetMemoByVideoId(id string) (*Memo, error)
}
