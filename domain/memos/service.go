package memos

import (
	"fmt"

	"github.com/Nexters/myply/domain/musics"
)

type Service interface {
	GetMemo(id string) (*Memo, error)
	AddMemo(videoID string, body string, deviceToken string) (*Memo, error)
	UpdateBody(id string, body string, deviceToken string) (*Memo, error)
}

type memoService struct {
	repository   Repository
	musicService musics.Service
}

func NewMemoService(r Repository, musicService musics.Service) Service {
	return &memoService{r, musicService}
}

func (s *memoService) GetMemo(id string) (*Memo, error) {
	m, err := s.repository.GetMemo(id)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (s *memoService) AddMemo(videoID string, body string, deviceToken string) (*Memo, error) {
	memo, err := s.repository.GetMemoByVideoID(videoID)
	if memo != nil {
		return nil, &AlreadyExistsError{Msg: fmt.Sprintf("memo with videoID already exists. videoID=%s", videoID)}
	}

	switch err.(type) {
	case *NotFoundError:
		music, musicErr := s.musicService.GetMusic(videoID)
		if musicErr != nil {
			return nil, musicErr
		}

		id, insertErr := s.repository.AddMemo(deviceToken, videoID, body, music.YoutubeTags)
		if insertErr != nil {
			return nil, insertErr
		}

		return s.GetMemo(id)
	default:
		return nil, err
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

	if err = s.repository.UpdateBody(id, body); err != nil {
		return nil, err
	}

	return s.GetMemo(id)
}
