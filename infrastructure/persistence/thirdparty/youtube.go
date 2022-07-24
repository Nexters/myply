package thirdparty

import (
	"github.com/Nexters/myply/domain/entity"
)

type Youtube interface {
	Search(query string) (*entity.Musics, error)
}
