package controller

import (
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

type RecommendResponse struct {
	BaseResponse
	Data RecommendResponseData `json:"data"`
}

func NewTagController(ts tag.TagService) TagController {
	return &tagController{service: ts}
}

// @Summary Get recommended tags
// @Description get tags recommended by myply
// @Tags tags
// @Accept json
// @Produce json
// @Success 200 {object} RecommendResponse
// @Failure 500
// @Router /tags/recommend [get]
// @Security ApiKeyAuth
func (tc *tagController) Recommend() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tags, err := tc.service.Recommend()

		if err != nil {
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
