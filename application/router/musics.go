package router

import (
	"github.com/Nexters/myply/application/controller"
	"github.com/gofiber/fiber/v2"
)

type MusicsRouter = Router

type musicsRouter struct {
	musicsController controller.MusicsController
}

func NewMusicsRouter(mc controller.MusicsController) MusicsRouter {
	return &musicsRouter{musicsController: mc}
}

func (mr *musicsRouter) Init(root *fiber.Router) {
	musicsRouter := (*root).Group("/musics")
	{
		musicsRouter.Get("/popular", mr.musicsController.GetPopularList())
	}
}
