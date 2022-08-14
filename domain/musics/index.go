package musics

import (
	"fmt"

	"github.com/google/wire"
)

var Set = wire.NewSet(NewMusicService)

func GenerateCacheKey(query, token string) string {
	return fmt.Sprintf("%s:%s", query, token)
}
