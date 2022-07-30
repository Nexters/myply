package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMusicsRouter)

type Router interface {
	Init(root *fiber.Router)
}
