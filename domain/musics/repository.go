package musics

type MusicRepository interface {
	GetMusicList(key string) (*Musics, bool, error)
	SaveMusicList(key string, musicList []byte) error
	GetPlayListBy(order string) (*Musics, error)
}
