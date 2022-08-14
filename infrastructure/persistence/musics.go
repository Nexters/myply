package persistence

import (
	"encoding/json"
	"fmt"

	"github.com/Nexters/myply/domain/musics"
	"github.com/Nexters/myply/infrastructure/clients"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/persistence/cache"
	v3 "google.golang.org/api/youtube/v3"
)

type MusicRepository struct {
	c  *configs.Config
	cc cache.Cache
	yc clients.YoutubeClient
}

func NewMusicRepository(c *configs.Config, cc cache.Cache, yc clients.YoutubeClient) musics.MusicRepository {
	return &MusicRepository{c, cc, yc}

}

func (m *MusicRepository) buildMusicListResponse(items []*v3.SearchResult, nextPageToken string) (*musics.MusicListDto, error) {
	ids := m.yc.ParseVideoIds(items)
	musicListResponse := make(musics.Musics, len(items))
	tags, err := m.yc.ParseVideoTags(ids)
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		switch item.Id.Kind {
		case clients.YoutubeVideoType:
			musicListResponse[i] = *musics.NewMusic(
				item.Id.VideoId,
				item.Snippet.Thumbnails.Default.Url,
				item.Snippet.Title,
				tags[item.Id.VideoId],
			)
		}
	}
	return &musics.MusicListDto{
		Musics: &musicListResponse, NextPageToken: nextPageToken}, nil
}

func (m *MusicRepository) GetMusicList(q, pageToken string) (musicListResponse *musics.MusicListDto, isCached bool, err error) {
	redisKey := musics.GenerateRedisKey(q, pageToken)
	data, _ := m.cc.Get(redisKey)
	if data != nil {
		if err = json.Unmarshal(data, &musicListResponse); err == nil {
			fmt.Printf("CACHE HIT of %s\n", redisKey)
			return musicListResponse, true, nil
		}
	}

	var musicsV3 *v3.SearchListResponse
	musicsV3, err = m.yc.SearchPlaylist(fmt.Sprintf("playlist,%s", q), "", pageToken)
	if err != nil {
		return nil, false, err
	}

	if musicListResponse, err = m.buildMusicListResponse(musicsV3.Items, musicsV3.NextPageToken); err != nil {
		return nil, false, err
	}

	return musicListResponse, false, nil
}

func (m *MusicRepository) SaveMusicList(key string, musicList []byte) error {
	return m.cc.Set(key, musicList, m.c.MongoCacheTTL)
}

func (m *MusicRepository) GetPlayListBy(order, pageToken string) (musicListResponse *musics.MusicListDto, err error) {
	var musicsV3 *v3.SearchListResponse
	musicsV3, err = m.yc.SearchPlaylist("", order, pageToken)
	if err != nil {
		return nil, err
	}

	if musicListResponse, err = m.buildMusicListResponse(musicsV3.Items, musicsV3.NextPageToken); err != nil {
		return nil, err
	}

	return musicListResponse, nil
}
