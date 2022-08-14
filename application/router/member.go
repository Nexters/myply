package router

import (
	"github.com/Nexters/myply/application/controller"
	"github.com/gofiber/fiber/v2"
)

type MemberRouter interface {
	Init(root *fiber.Router)
}

type memberRouter struct {
	controller controller.MemberController
}

func NewMemberRouter(mc controller.MemberController) MemberRouter {
	return &memberRouter{controller: mc}
}

func (mr *memberRouter) Init(root *fiber.Router) {
	memberRouter := (*root).Group("/members")
	{
		memberRouter.Get("/", mr.controller.Get())
		memberRouter.Post("/", mr.controller.SignUp())
		memberRouter.Patch("/", mr.controller.Update())
	}
}
