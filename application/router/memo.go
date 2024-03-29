package router

import (
	"github.com/Nexters/myply/application/controller"
	"github.com/gofiber/fiber/v2"
)

type MemoRouter interface {
	Init(root *fiber.Router)
}

type memoRouter struct {
	controller controller.MemoController
}

func NewMemoRouter(mc controller.MemoController) MemoRouter {
	return &memoRouter{controller: mc}
}

func (mr *memoRouter) Init(root *fiber.Router) {
	memoRouter := (*root).Group("/memos")
	{
		memoRouter.Get("/", mr.controller.GetMemos)
		memoRouter.Get("/:memoID", mr.controller.GetMemo)
		memoRouter.Post("/", mr.controller.AddMemo)
		memoRouter.Patch("/:memoID", mr.controller.UpdateMemo)
		memoRouter.Delete("/:memoIdOrYoutubeID", mr.controller.DeleteMemo)
	}

}
