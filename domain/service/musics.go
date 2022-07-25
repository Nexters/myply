package service

import (
	"github.com/Nexters/myply/domain/entity"
	"github.com/Nexters/myply/domain/repository"
)

type MusicsService interface {
	GetPopularList() (*entity.Musics, error) // TODO: pagination
}

type musicsService struct {
	musicsRepository repository.MusicsRepository
}

func NewMusicsService(mr repository.MusicsRepository) MusicsService {
	return &musicsService{musicsRepository: mr}
}

func (ms *musicsService) GetPopularList() (*entity.Musics, error) {
	musics, err := ms.musicsRepository.GetPopularList()

	if err != nil {
		return nil, err
	}

	return musics, err
}
