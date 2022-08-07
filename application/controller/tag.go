package controller

import (
	"fmt"

	"github.com/Nexters/myply/domain/tag"
	"github.com/gofiber/fiber/v2"
)

type TagController interface {
	Recommend() fiber.Handler
}

type tagController struct {
	service tag.TagService
}

type RecommendResponseData struct {
	Tags []string `json:"tags"`
}

func NewTagController(ts tag.TagService) TagController {
	return &tagController{service: ts}
}

func (tc *tagController) Recommend() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tags, err := tc.service.Recommend()

		if err != nil {
			fmt.Println(err.Error())

			return c.Status(500).JSON(BaseResponse{
				Code:    500,
				Message: "fail: internal server error",
				Data:    nil,
			})
		}

		data := RecommendResponseData{Tags: tags.Labels()}
		res := BaseResponse{
			Code:    200,
			Message: "success: get tags",
			Data:    data,
		}

		return c.Status(200).JSON(res)
	}
}
