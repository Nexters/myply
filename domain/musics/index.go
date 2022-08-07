package musics

import "github.com/google/wire"

var Set = wire.NewSet(NewMusicService)
