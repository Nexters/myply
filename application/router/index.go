package router

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewMemberRouter,
	NewMemoRouter,
	NewMusicsRouter,
	NewTagRouter,
)
