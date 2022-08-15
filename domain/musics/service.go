package musics

import (
	"encoding/json"

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
	GetMusicList(query string, pageToken string) (*MusicListDto, error)
	GetPlayListBy(order Order, pageToken string) (*MusicListDto, error)
	GetTags(videoId string) ([]string, error)
}

type MusicListDto struct {
	Musics        *Musics
	NextPageToken string // optional
}

type musicService struct {
	logger          *zap.SugaredLogger
	musicRepository MusicRepository
}

func NewMusicService(l *zap.SugaredLogger, mr MusicRepository) Service {
	return &musicService{logger: l, musicRepository: mr}
}

func (ms *musicService) GetPlayListBy(order Order, pageToken string) (*MusicListDto, error) {
	return ms.musicRepository.GetPlayListBy(order.String(), pageToken)
}

func (ms *musicService) GetMusicList(query string, pageToken string) (*MusicListDto, error) {
	musics, isCached, err := ms.musicRepository.GetMusicList(query, pageToken)
	if err != nil {
		return nil, err
	}

	if !isCached {
		musicsBytes, err := json.Marshal(musics)
		if err != nil {
			return nil, err
		}

		err = ms.musicRepository.SaveMusicList(GenerateCacheKey(query, pageToken), musicsBytes)
		if err != nil {
			return nil, err
		}
		return musics, nil
	}
	return musics, nil
}

func (ms *musicService) GetTags(videoId string) ([]string, error) {
	return ms.musicRepository.GetTags(videoId)
}
