package router

import (
	"github.com/Nexters/myply/application/controller"
	"github.com/gofiber/fiber/v2"
)

func SetMusicsRouter(root *fiber.Router, mc controller.MusicsController) {
	musicsRouter := (*root).Group("/musics")
	{
		musicsRouter.Get("/popular", mc.GetPopularList())
	}
}
