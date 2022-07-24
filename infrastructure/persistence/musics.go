package persistence

import (
	"github.com/Nexters/myply/domain/entity"
	"github.com/Nexters/myply/domain/repository"
	"github.com/Nexters/myply/infrastructure/persistence/thirdparty"
)

type MusicsRepository struct {
	youtube thirdparty.Youtube
}

func NewMusicRepository(ys thirdparty.Youtube) repository.MusicsRepository {
	return &MusicsRepository{youtube: ys}
}

func (ms *MusicsRepository) GetPopularList() (*entity.Musics, error) {
	// TODO: save redis
	musics, err := ms.youtube.Search("playlist")

	if err != nil {
		return nil, err
	}

	return musics, nil
}
