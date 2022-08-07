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

func (m *MusicRepository) buildMusicListResponse(items []*v3.SearchResult) (*musics.Musics, error) {
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
	return &musicListResponse, nil
}

func (m *MusicRepository) GetMusicList(q string) (musicListResponse *musics.Musics, isCached bool, err error) {
	data, _ := m.cc.Get(q)
	if data != nil {
		if err = json.Unmarshal(data, musicListResponse); err == nil {
			return musicListResponse, true, nil
		}

		fmt.Println(err.Error())
	}

	var musicsV3 *v3.SearchListResponse
	musicsV3, err = m.yc.SearchPlaylist(fmt.Sprintf("playlist,%s", q))
	if err != nil {
		return nil, false, err
	}

	if musicListResponse, err = m.buildMusicListResponse(musicsV3.Items); err != nil {
		return nil, false, err
	}

	return musicListResponse, false, nil
}

func (m *MusicRepository) SaveMusicList(key string, musicList []byte) error {
	return m.cc.Set(key, musicList, m.c.MongoCacheTTL)
}
