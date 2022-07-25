package repository

import "github.com/Nexters/myply/domain/entity"

type MusicsRepository interface {
	GetPopularList() (*entity.Musics, error) // TODO: pagination
}
