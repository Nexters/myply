package controller

import (
	"strings"

	"github.com/Nexters/myply/domain/member"
	"github.com/gofiber/fiber/v2"
)

type MemberController interface {
	SignUp() fiber.Handler
	Get() fiber.Handler
	Update() fiber.Handler
}

type signUpDTO struct {
	DeviceToken string   `json:"deviceToken"`
	Name        string   `json:"name"`
	Keywords    []string `json:"keywords"`
}

type updateDTO struct {
	Name     *string  `json:"name"`
	Keywords []string `json:"keywords"`
}

type memberResponseData struct {
	DeviceToken string   `json:"deviceToken"`
	Name        string   `json:"name"`
	Keywords    []string `json:"keywords"`
}

type MemberResponse struct {
	BaseResponse
	Data memberResponseData `json:"data"`
}

type memberController struct {
	service member.MemberService
}

func NewMemberController(ms member.MemberService) MemberController {
	return &memberController{service: ms}
}

// @Summary Sign up
// @Description 회원가입
// @Tags members
// @Accept json
// @Produce json
// @Param body body signUpDTO true "sign up body"
// @Success 200 {object} BaseResponse
// @Failure 409 "Account already exist"
// @Failure 500 "Internal server error"
// @Router /members/ [post]
func (mc *memberController) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto := new(signUpDTO)

		if err := c.BodyParser(dto); err != nil {
			return err
		}

		signUpErr := mc.service.SignUp(
			dto.DeviceToken,
			dto.Name,
			dto.Keywords,
		)
		if signUpErr != nil {
			msg := signUpErr.Error()
			if strings.HasPrefix(msg, "409:") {
				return c.Status(409).JSON(BaseResponse{
					Code:    409,
					Message: "fail: account already exist",
					Data:    nil,
				})
			}

			return c.Status(500).JSON(BaseResponse{
				Code:    500,
				Message: "fail: internal server error",
				Data:    nil,
			})
		}

		return c.Status(201).JSON(BaseResponse{
			Code:    201,
			Message: "success: sign up",
			Data:    nil,
		})
	}
}

// @Summary Get user info
// @Description 내 상세정보를 얻는다.
// @Description - Device-Token 헤더값이 필요하다.
// @Tags members
// @Accept json
// @Produce json
// @Success 200 {object} MemberResponse
// @Failure 401 "Unautorized"
// @Failure 500 "Internal server error"
// @Router /members/ [get]
// @Security ApiKeyAuth
func (mc *memberController) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		member := c.Locals("member").(*member.Member)

		data := memberResponseData{
			DeviceToken: member.DeviceToken,
			Name:        member.Name,
			Keywords:    member.Keywords,
		}
		res := BaseResponse{
			Code:    200,
			Message: "Success",
			Data:    data,
		}

		return c.Status(200).JSON(res)
	}
}

// @Summary Update name or keywords
// @Description 내 정보를 업데이트 한다.
// @Description - Device-Token 헤더값이 필요하다.
// @Description - 이름만 업데이트 할경우 "name" 필드만, 키워드만 업데이트 할 경우 "keywords" 필드만 넘겨주면 된다.
// @Tags members
// @Accept json
// @Produce json
// @Param body body updateDTO true "update body"
// @Success 200 {object} BaseResponse
// @Failure 500 "Internal server error"
// @Router /members/ [patch]
func (mc *memberController) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto := new(updateDTO)
		if err := c.BodyParser(dto); err != nil {
			return err
		}

		member := c.Locals("member").(*member.Member)
		deviceToken := member.DeviceToken

		if err := mc.service.Update(deviceToken, dto.Name, dto.Keywords); err != nil {
			return err
		}

		return c.Status(200).JSON(BaseResponse{
			Code:    200,
			Message: "Success",
			Data:    nil,
		})
	}
}
