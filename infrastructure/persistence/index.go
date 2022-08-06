package persistence

import (
	"github.com/Nexters/myply/infrastructure/persistence/cache"
	"github.com/Nexters/myply/infrastructure/persistence/db"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewMusicRepository,
	NewMemoRepository,
	NewMemberRepository,
	NewTagRepository,
	db.NewMongoDB,
	cache.NewMongoCacheDB,
)
