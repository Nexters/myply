package musics

type MusicRepository interface {
	GetMusic(videoId string) (*Music, error)
	GetMusicList(key, pageToken string) (*MusicListDto, bool, error)
	SaveMusicList(key string, musicList []byte) error
	GetPlayListBy(order, pageToken string) (*MusicListDto, error)
}
