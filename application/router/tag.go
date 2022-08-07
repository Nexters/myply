package router

import (
	"github.com/Nexters/myply/application/controller"
	"github.com/gofiber/fiber/v2"
)

type TagRouter interface {
	Init(root *fiber.Router)
}

type tagRouter struct {
	controller controller.TagController
}

func NewTagRouter(tc controller.TagController) TagRouter {
	return &tagRouter{controller: tc}
}

func (tr *tagRouter) Init(root *fiber.Router) {
	tagRouter := (*root).Group("/tags")
	{
		tagRouter.Get("/recommend", tr.controller.Recommend())
	}
}
