package persistence

import (
	"github.com/Nexters/myply/domain/repository"
)

type MusicRepository struct {
}

func NewMusicRepository() repository.MusicRepository {
	return &MusicRepository{}
}
