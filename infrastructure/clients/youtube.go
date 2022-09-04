package clients

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nexters/myply/domain/musics"
	"github.com/Nexters/myply/infrastructure/configs"
	"google.golang.org/api/option"
	v3 "google.golang.org/api/youtube/v3"
)

var (
	targetTypes    = []string{"video"} // video, channel, playlist
	emptyVideoInfo = VideoInfo{}
)

const (
	YoutubeVideoType   = "youtube#video"
	YoutubeChannelType = "youtube#channel"
)

const (
	videoCategory = "10" // music
	regionCode    = "KR"
	videoDuration = "long" // more than 20m
	maxResults    = 25
	defaultOrder  = "relevance"
	defaultQuery  = "playlist,"
)

type TagMap map[string][]string

type VideoInfo struct {
	VideoID      string
	ThumbnailURL string
	Title        string
	Tags         []string
}

type VideoInfos []VideoInfo

type YoutubeClient interface {
	GetMusicDetail(videoID string) (*VideoInfo, error)
	GetMusics(videoIDs []string) (VideoInfos, error)
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

func (v *VideoInfo) ToEntity() *musics.Music {
	if v == &emptyVideoInfo {
		return &musics.EmptyMusic
	}
	return &musics.Music{
		YoutubeVideoID: v.VideoID,
		ThumbnailURL:   v.ThumbnailURL,
		Title:          v.Title,
		YoutubeTags:    v.Tags,
	}
}

func (vs *VideoInfos) ToEntity() musics.Musics {
	ms := musics.Musics{}
	for _, v := range *vs {
		ms = append(ms, *v.ToEntity())
	}
	return ms
}

func (yc *youtubeClient) GetMusics(videoIDs []string) (VideoInfos, error) {
	// The id parameter value is a comma-separated list of YouTube video IDs
	call := yc.service.Videos.List([]string{"snippet"}).Id(strings.Join(videoIDs, ","))
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	if len(response.Items) <= 0 {
		return nil, &musics.NotFoundError{Msg: fmt.Sprintf("youtube music is not found. ids=%s", videoIDs)}
	}

	result := VideoInfos{}
	for i, v := range response.Items {
		if v.Id == videoIDs[i] {
			result = append(result, VideoInfo{
				VideoID:      v.Id,
				ThumbnailURL: v.Snippet.Thumbnails.Default.Url,
				Title:        v.Snippet.Title,
				Tags:         v.Snippet.Tags,
			})
		} else {
			result = append(result, emptyVideoInfo)
		}
	}
	return result, nil
}

func (yc *youtubeClient) GetMusicDetail(videoID string) (*VideoInfo, error) {
	call := yc.service.Videos.List([]string{"snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	if len(response.Items) <= 0 {
		return nil, &musics.NotFoundError{Msg: fmt.Sprintf("youtube music is not found. id=%s", videoID)}
	}

	result := response.Items[0]
	return &VideoInfo{
		VideoID:      videoID,
		ThumbnailURL: result.Snippet.Thumbnails.Default.Url,
		Title:        result.Snippet.Title,
		Tags:         result.Snippet.Tags,
	}, nil
}

func (yc *youtubeClient) SearchPlaylist(q, order, pageToken string) (*v3.SearchListResponse, error) {
	if order == "" {
		order = defaultOrder
	}
	if q == "" {
		q = defaultQuery
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
