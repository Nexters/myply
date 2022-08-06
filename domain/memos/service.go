package memos

import (
	"errors"
)

type Service interface {
	GetMemo(id string) (*Memo, error)
	AddMemo(videoId string, body string, deviceToken string) (memoId string, e error)
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
		return "", AlreadyExistsException
	}

	if !(errors.Is(err, NotFoundException)) {
		return "", err
	}

	return (*s.repository).SaveMemo(videoId, body, deviceToken)
}
