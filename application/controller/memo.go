package controller

import (
	"github.com/Nexters/myply/domain/memos"
	"github.com/gofiber/fiber/v2"
)

type MemoController interface {
	GetMemo(ctx *fiber.Ctx) error
	AddMemo(ctx *fiber.Ctx) error
	UpdateMemo(ctx *fiber.Ctx) error
}

type memoController struct {
	service *memos.Service
}

func NewMemoController(s *memos.Service) MemoController {
	return &memoController{service: s}
}

func (c *memoController) GetMemo(ctx *fiber.Ctx) error {
	var resp Response

	id := ctx.Params("memoID")

	m, err := (*c.service).GetMemo(id)
	if err != nil {
		return c.handleErrors(err)
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

	token, err := c.deviceToken(ctx)
	if err != nil {
		return err
	}

	id, err := (*c.service).AddMemo(req.YoutubeVideoId, req.Body, token)
	if err != nil {
		return c.handleErrors(err)
	}

	memoResp := MemoResponse{MemoId: id, ThumbnailURL: "", Title: "", Body: "", Keywords: []string{}}
	resp = Response{
		code:    fiber.StatusCreated,
		message: "",
		data:    memoResp.toMap(),
	}

	return ctx.Status(fiber.StatusCreated).JSON(resp.toMap())
}

func (c *memoController) UpdateMemo(ctx *fiber.Ctx) error {
	var resp Response

	id := ctx.Params("memoID")

	req := new(patchRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	token, err := c.deviceToken(ctx)
	if err != nil {
		return err
	}

	m, err := (*c.service).UpdateBody(id, req.Body, token)
	if err != nil {
		return c.handleErrors(err)
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

func (c *memoController) deviceToken(ctx *fiber.Ctx) (string, error) {
	token := ctx.GetReqHeaders()["Device-Token"]
	if token == "" {
		return "", fiber.NewError(fiber.StatusBadRequest, "empty device-token")
	}

	return token, nil
}

func (c *memoController) handleErrors(err error) error {
	switch err.(type) {
	case *memos.NotFoundError:
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	case *memos.AlreadyExistsError, *memos.IllegalDeviceTokenError:
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	default:
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}

type addRequest struct {
	YoutubeVideoId string `json:"youtubeVideoId"`
	Body           string `json:"body"`
}

type patchRequest struct {
	Body string `json:"body"`
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
