package memos

import (
	"fmt"
)

type Service interface {
	GetMemo(id string) (*Memo, error)
	AddMemo(videoId string, body string, deviceToken string) (memoId string, e error)
	UpdateBody(id string, body string, deviceToken string) (*Memo, error)
}

type memoService struct {
	repository *Repository
}

func NewMemoService(r *Repository) *Service {
	var service Service
	service = &memoService{r}
	return &service
}

func (s *memoService) GetMemo(id string) (*Memo, error) {
	m, err := (*s.repository).GetMemo(id)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (s *memoService) AddMemo(videoId string, body string, deviceToken string) (memoId string, e error) {
	memo, err := (*s.repository).GetMemoByVideoId(videoId)
	if memo != nil {
		return "", &AlreadyExistsError{Msg: fmt.Sprintf("memo with videoID already exists. videoID=%s", videoId)}
	}

	switch err.(type) {
	case *NotFoundError:
		return (*s.repository).AddMemo(deviceToken, videoId, body, nil)
	default:
		return "", err
	}
}

func (s *memoService) UpdateBody(id string, body string, deviceToken string) (*Memo, error) {
	old, err := s.GetMemo(id)
	if err != nil {
		return nil, err
	}

	if old.DeviceToken != deviceToken {
		return nil, &IllegalDeviceTokenError{Msg: fmt.Sprintf("failed to update due to invalid device token")}
	}

	if err = (*s.repository).UpdateBody(id, body); err != nil {
		return nil, err
	}

	return s.GetMemo(id)
}

func (s *memoService) UpdateBody(id string, body string, deviceToken string) (*Memo, error) {
	m, err := (*s.repository).UpdateBody(id, body)
	if err != nil {
		return nil, err
	}

	return m, err
}
