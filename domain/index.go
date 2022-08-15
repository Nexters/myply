package domain

import (
	"github.com/Nexters/myply/domain/member"
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/domain/musics"
	"github.com/Nexters/myply/domain/tag"
	"github.com/google/wire"
)

var Set = wire.NewSet(musics.NewMusicService, memos.NewMemoService, member.NewMemberService, tag.NewTagService)
