package controller

import (
	"errors"
	"github.com/Nexters/myply/domain/memos"
	"github.com/gofiber/fiber/v2"
)

type MemoController interface {
	GetMemo(ctx *fiber.Ctx) error
	AddMemo(ctx *fiber.Ctx) error
}

type memoController struct {
	service *memos.Service
}

func NewMemoController(s *memos.Service) *MemoController {
	var c MemoController
	c = &memoController{service: s}
	return &c
}

func (c *memoController) GetMemo(ctx *fiber.Ctx) error {
	var resp Response

	id := ctx.Params("id")

	m, err := (*c.service).GetMemo(id)
	if err != nil {
		switch {
		case errors.Is(err, memos.NotFoundException):
			resp = Response{code: fiber.StatusNotFound, message: err.Error(), data: fiber.Map{}}
			return ctx.Status(fiber.StatusNotFound).JSON(resp.toMap())
		default:
			resp = Response{code: fiber.StatusInternalServerError, message: err.Error(), data: fiber.Map{}}
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.toMap())
		}
	}

	// TODO: respond real data
	memoResp := MemoResponse{MemoId: m.Id, ThumbnailURL: "", Title: "", Body: m.Body, Keywords: []string{}}
	resp = Response{
		code:    fiber.StatusOK,
		message: "",
		data:    memoResp.toMap(),
	}

	return ctx.Status(fiber.StatusOK).JSON(resp.toMap())
}

func (c *memoController) AddMemo(ctx *fiber.Ctx) error {
	var resp Response

	req := new(addRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	token := ctx.GetReqHeaders()["Device-Token"]
	if token == "" {
		resp = Response{
			code:    fiber.StatusBadRequest,
			message: "failed due to empty device-token",
			data:    fiber.Map{},
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(resp.toMap())
	}

	id, err := (*c.service).AddMemo(req.YoutubeVideoId, req.Body, token)
	if err != nil {
		resp = Response{message: err.Error(), data: fiber.Map{}}
		switch {
		case errors.Is(err, memos.AlreadyExistsException):
			resp.code = fiber.StatusBadRequest
			return ctx.Status(fiber.StatusBadRequest).JSON(resp.toMap())
		default:
			resp.code = fiber.StatusInternalServerError
			return ctx.Status(fiber.StatusInternalServerError).JSON(resp.toMap())
		}
	}

	memoResp := MemoResponse{MemoId: id, ThumbnailURL: "", Title: "", Body: "", Keywords: []string{}}
	resp = Response{
		code:    fiber.StatusCreated,
		message: "",
		data:    memoResp.toMap(),
	}

	return ctx.Status(fiber.StatusCreated).JSON(resp.toMap())

}

type addRequest struct {
	YoutubeVideoId string `json:"youtubeVideoId"`
	Body           string `json:"body"`
}

type MemoResponse struct {
	MemoId       string   `json:"memoID"`
	ThumbnailURL string   `json:"thumbnailURL"`
	Title        string   `json:"title"`
	Body         string   `json:"body"`
	Keywords     []string `json:"keywords"`
}

func (r *MemoResponse) toMap() fiber.Map {
	return fiber.Map{
		"memoId":       r.MemoId,
		"thumbnailURL": r.ThumbnailURL,
		"title":        r.Title,
		"body":         r.Body,
		"keywords":     r.Keywords,
	}
}
