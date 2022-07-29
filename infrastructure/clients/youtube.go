package clients

import (
	"context"

	"github.com/Nexters/myply/domain/entity"
	"github.com/Nexters/myply/infrastructure/configs"
	"google.golang.org/api/option"
	v3 "google.golang.org/api/youtube/v3"
)

var (
	targetTypes = []string{"video"} // video, channel, playlist
)

const (
	videoCategory = "10" // music
	regionCode    = "kr"
	videoDuration = "long" // more than 20m
	maxResults    = 25

	youtubeVideoType   = "youtube#video"
	youtubeChannelType = "youtube#channel"
)

type tagMap map[string][]string

type YoutubeClient interface {
	SearchPlaylist(q string) (*entity.Musics, error)
}

type youtubeClient struct {
	service *v3.Service
}

func NewYoutubeClient(config *configs.Config) (YoutubeClient, error) {
	// TODO: if needed use option.WithGRPCConnectionPool()
	service, err := v3.NewService(
		context.Background(),
		option.WithAPIKey(config.YoutubeAPIKey),
	)

	if err != nil {
		return nil, err
	}
	return &youtubeClient{service}, nil
}

func (yc *youtubeClient) SearchPlaylist(q string) (*entity.Musics, error) {
	call := yc.service.Search.List([]string{"id, snippet"}).
		Q(q).
		Type(targetTypes...).
		RegionCode(regionCode).
		VideoCategoryId(videoCategory).
		VideoDuration(videoDuration).
		// Order(). // TODO: viewCount, date
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	return yc.buildMusicListResponse(response.Items)
}

func (yc *youtubeClient) getVideoTags(ids []string) (tagMap, error) {
	tags := make(tagMap)

	call := yc.service.Videos.List([]string{"snippet"}).Id(ids...)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	for _, i := range response.Items {
		tags[i.Id] = i.Snippet.Tags
	}
	return tags, nil
}

func getVideoIds(items []*v3.SearchResult) []string {
	ids := make([]string, len(items))
	for _, item := range items {
		ids = append(ids, item.Id.VideoId)
	}
	return ids
}

func (yc *youtubeClient) buildMusicListResponse(items []*v3.SearchResult) (*entity.Musics, error) {
	ids := getVideoIds(items)
	musics := make(entity.Musics, len(items))
	tags, err := yc.getVideoTags(ids)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		switch item.Id.Kind {
		case youtubeVideoType:
			musics[i] = *entity.NewMusic(
				item.Id.VideoId,
				item.Snippet.Thumbnails.Default.Url,
				item.Snippet.Title,
				tags[item.Id.VideoId],
			)
		}
	}
	return &musics, nil
}
