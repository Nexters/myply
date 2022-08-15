package musics

type MusicRepository interface {
	GetMusicList(key, pageToken string) (*MusicListDto, bool, error)
	SaveMusicList(key string, musicList []byte) error
	GetPlayListBy(order, pageToken string) (*MusicListDto, error)
	GetTags(videoId string) ([]string, error)
}
