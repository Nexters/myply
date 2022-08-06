package memos

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMemoService)
