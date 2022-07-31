package domain

import (
	"github.com/Nexters/myply/domain/memos"
	"github.com/google/wire"
)

var Set = wire.NewSet(memos.NewMemoService)
