package controller

import (
	"github.com/Nexters/myply/domain/service"
	"github.com/gofiber/fiber/v2"
)

type MusicsController interface {
	GetPopularList() fiber.Handler // TODO: pagination
}

type musicsController struct {
	musicsService service.MusicsService
}

func NewMusicsController(ms service.MusicsService) MusicsController {
	return &musicsController{musicsService: ms}
}

func (mc *musicsController) GetPopularList() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		musics, err := mc.musicsService.GetPopularList()

		if err != nil {
			return nil
		}

		return ctx.JSON(musics)
	}
}
