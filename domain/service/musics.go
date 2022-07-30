package service

import (
	"fmt"
	"strings"

	"github.com/Nexters/myply/domain/entity"
	"github.com/Nexters/myply/domain/repository"
	"github.com/Nexters/myply/infrastructure/clients"
	"go.uber.org/zap"
)

type MusicsService interface {
	GetPopularList(rawQueries []string) (*entity.Musics, error) // TODO: pagination
}

type musicService struct {
	logger          *zap.SugaredLogger
	youtubeClient   clients.YoutubeClient
	musicRepository repository.MusicRepository
}

func NewMusicService(l *zap.SugaredLogger, yc clients.YoutubeClient, mr repository.MusicRepository) MusicsService {
	return &musicService{logger: l, youtubeClient: yc, musicRepository: mr}
}

func (ms *musicService) GetPopularList(rawQueries []string) (*entity.Musics, error) {
	queries := fmt.Sprintf("playlist,%s", strings.Join(rawQueries, ","))
	ms.logger.Infof("q=%s", queries)
	musics, err := ms.youtubeClient.SearchPlaylist(queries)

	// TODO: cache to redis
	if err != nil {
		return nil, err
	}
	return musics, nil
}
