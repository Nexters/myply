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

// @Summary      Retrieve Memo
// @Description  메모 조회
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param memoID path string true "memoID to retrieve"
// @Success      200  {object}   MemoResponse
// @Failure      404
// @Failure      500
// @Router       /memos/{memoID} [get]
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

// @Summary      Add Memo
// @Description  메모 생성
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param Body body AddRequest true "memo request body"
// @Success      200  {object}   MemoResponse
// @Failure      500
// @Router       /memos/ [post]
func (c *memoController) AddMemo(ctx *fiber.Ctx) error {
	var resp Response

	req := new(AddRequest)
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

// @Summary      Update Memo
// @Description  메모 수정
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param Body body PatchRequest true "memo request body"
// @Success      200  {object}   MemoResponse
// @Failure      500
// @Router       /memos/ [patch]
func (c *memoController) UpdateMemo(ctx *fiber.Ctx) error {
	var resp Response

	id := ctx.Params("memoID")

	req := new(PatchRequest)
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

type AddRequest struct {
	YoutubeVideoId string `json:"youtubeVideoId"`
	Body           string `json:"body"`
}

type PatchRequest struct {
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
