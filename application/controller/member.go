package controller

import (
	"strings"

	"github.com/Nexters/myply/domain/member"
	"github.com/gofiber/fiber/v2"
)

type MemberController interface {
	SignUp() fiber.Handler
}

type signUpDTO struct {
	DeviceToken string `json:"deviceToken"`
	Name        string `json:"name"`
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

		signUpErr := mc.service.SignUp(dto.DeviceToken, dto.Name)
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
