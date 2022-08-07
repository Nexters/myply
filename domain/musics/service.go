package musics

import (
	"encoding/json"
	"strings"

	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/persistence/cache"
	"go.uber.org/zap"
)

type Service interface {
	GetPopularList(rawQueries []string) (*Musics, error) // TODO: pagination
}

type musicService struct {
	logger          *zap.SugaredLogger
	musicRepository MusicRepository
}

func NewMusicService(l *zap.SugaredLogger, mr MusicRepository, c *configs.Config, cc cache.Cache) Service {
	return &musicService{logger: l, musicRepository: mr}
}

// MongoCacheTTL
func (ms *musicService) GetPopularList(rawQueries []string) (*Musics, error) {
	// TODO: parse korean sentence to []noun
	queries := strings.Join(rawQueries, ",")
	musics, isCached, err := ms.musicRepository.GetMusicList(queries)
	if err != nil {
		return nil, err
	}

	if !isCached {
		musicsBytes, err := json.Marshal(musics)
		if err != nil {
			return nil, err
		}
		err = ms.musicRepository.SaveMusicList(queries, musicsBytes)
		if err != nil {
			return nil, err
		}
		return musics, nil
	}
	ms.logger.Infof("\n[Cache hit]\n- key: %s\n- len of musics: %d\n", queries, len(*musics))
	return musics, nil
}
