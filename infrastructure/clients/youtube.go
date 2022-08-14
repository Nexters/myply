package clients

import (
	"context"

	"github.com/Nexters/myply/infrastructure/configs"
	"google.golang.org/api/option"
	v3 "google.golang.org/api/youtube/v3"
)

var (
	targetTypes = []string{"video"} // video, channel, playlist
)

const (
	YoutubeVideoType   = "youtube#video"
	YoutubeChannelType = "youtube#channel"
)

const (
	videoCategory = "10" // music
	regionCode    = "kr"
	videoDuration = "long" // more than 20m
	maxResults    = 25
	defaultOrder  = "relevance"
)

type TagMap map[string][]string

type YoutubeClient interface {
	SearchPlaylist(q string, order string, pageToken string) (*v3.SearchListResponse, error)
	ParseVideoIds(items []*v3.SearchResult) []string
	ParseVideoTags(ids []string) (TagMap, error)
}

type youtubeClient struct {
	service *v3.Service
}

func NewYoutubeClient(config *configs.Config) (YoutubeClient, error) {
	// if needed use option.WithGRPCConnectionPool()
	service, err := v3.NewService(
		context.Background(),
		option.WithAPIKey(config.YoutubeAPIKey),
	)

	if err != nil {
		return nil, err
	}
	return &youtubeClient{service}, nil
}

func (yc *youtubeClient) SearchPlaylist(q, order, pageToken string) (*v3.SearchListResponse, error) {
	if order == "" {
		order = defaultOrder
	}

	call := yc.service.Search.List([]string{"id, snippet"}).
		Q(q).
		Type(targetTypes...).
		RegionCode(regionCode).
		PageToken(pageToken). // When pageToken == "", token will be ignored.
		VideoCategoryId(videoCategory).
		VideoDuration(videoDuration).
		Order(order).
		MaxResults(maxResults)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (yc *youtubeClient) ParseVideoIds(items []*v3.SearchResult) []string {
	ids := make([]string, len(items))
	for _, item := range items {
		ids = append(ids, item.Id.VideoId)
	}
	return ids
}

func (yc *youtubeClient) ParseVideoTags(ids []string) (TagMap, error) {
	tags := make(TagMap)

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
