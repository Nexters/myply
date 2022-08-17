package controller

import (
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/domain/musics"
	"github.com/gofiber/fiber/v2"
)

type MemoController interface {
	GetMemo(ctx *fiber.Ctx) error
	AddMemo(ctx *fiber.Ctx) error
	UpdateMemo(ctx *fiber.Ctx) error
}

type memoController struct {
	memoService  memos.Service
	musicService musics.Service
}

func NewMemoController(memoService memos.Service, musicService musics.Service) MemoController {
	return &memoController{memoService: memoService, musicService: musicService}
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
// @Security ApiKeyAuth
func (c *memoController) GetMemo(ctx *fiber.Ctx) error {
	id := ctx.Params("memoID")

	m, err := c.memoService.GetMemo(id)
	if err != nil {
		return c.handleErrors(err)
	}

	resp, err := c.generateResponse(m, fiber.StatusOK)
	if err != nil {
		return c.handleErrors(err)
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
// @Security ApiKeyAuth
func (c *memoController) AddMemo(ctx *fiber.Ctx) error {
	req := new(AddRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	token, err := c.deviceToken(ctx)
	if err != nil {
		return err
	}

	m, err := c.memoService.AddMemo(req.YoutubeVideoID, req.Body, token)
	if err != nil {
		return c.handleErrors(err)
	}

	resp, err := c.generateResponse(m, fiber.StatusCreated)
	if err != nil {
		return c.handleErrors(err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(resp.toMap())
}

// @Summary      Update Memo
// @Description  메모 수정
// @Tags         memos
// @Accept       json
// @Produce      json
// @Param memoID path string true "memoID to retrieve"
// @Param Body body PatchRequest true "memo request body"
// @Success      200  {object}   MemoResponse
// @Failure      500
// @Router       /memos/{memoID} [patch]
// @Security ApiKeyAuth
func (c *memoController) UpdateMemo(ctx *fiber.Ctx) error {
	id := ctx.Params("memoID")

	req := new(PatchRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	token, err := c.deviceToken(ctx)
	if err != nil {
		return err
	}

	m, err := c.memoService.UpdateBody(id, req.Body, token)
	if err != nil {
		return c.handleErrors(err)
	}

	resp, err := c.generateResponse(m, fiber.StatusOK)
	if err != nil {
		return c.handleErrors(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(resp.toMap())
}

func (c *memoController) generateResponse(memo *memos.Memo, successStatus int32) (Response, error) {
	music, err := c.musicService.GetMusic(memo.YoutubeVideoID)
	if err != nil {
		return Response{}, err
	}

	memoResp := MemoResponse{
		MemoID:       memo.ID,
		ThumbnailURL: music.ThumbnailURL,
		Title:        music.Title,
		Body:         memo.Body,
		Keywords:     music.YoutubeTags,
	}
	return Response{
		code:    successStatus,
		message: "",
		data:    memoResp.toMap(),
	}, nil
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
	YoutubeVideoID string `json:"youtubeVideoID"`
	Body           string `json:"body"`
}

type PatchRequest struct {
	Body string `json:"body"`
}

type MemoResponse struct {
	MemoID       string   `json:"memoID"`
	ThumbnailURL string   `json:"thumbnailURL"`
	Title        string   `json:"title"`
	Body         string   `json:"body"`
	Keywords     []string `json:"keywords"`
}

func (r *MemoResponse) toMap() fiber.Map {
	return fiber.Map{
		"memoId":       r.MemoID,
		"thumbnailURL": r.ThumbnailURL,
		"title":        r.Title,
		"body":         r.Body,
		"keywords":     r.Keywords,
	}
}
