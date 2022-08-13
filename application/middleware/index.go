package middleware

import "github.com/google/wire"

var Set = wire.NewSet(NewAuthMiddleware)
