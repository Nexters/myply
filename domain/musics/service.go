package musics

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"
)

type Order string

const (
	OrderByViewCount Order = "viewCount"
	OrderByDate            = "date"
	OrderByDefault         = "relevance"
)

func (o Order) String() string {
	switch o {
	case OrderByViewCount:
		return "viewCount"
	case OrderByDate:
		return "date"
	default:
		return "relevance"
	}
}

type Service interface {
	GetMusicList(rawQueries []string) (*Musics, error) // TODO: pagination
	GetPlayListBy(order Order) (*Musics, error)
}

type musicService struct {
	logger          *zap.SugaredLogger
	musicRepository MusicRepository
}

func NewMusicService(l *zap.SugaredLogger, mr MusicRepository) Service {
	return &musicService{logger: l, musicRepository: mr}
}

func (ms *musicService) GetPlayListBy(order Order) (*Musics, error) {
	return ms.musicRepository.GetPlayListBy(order.String())
}

// MongoCacheTTL
func (ms *musicService) GetMusicList(rawQueries []string) (*Musics, error) {
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
