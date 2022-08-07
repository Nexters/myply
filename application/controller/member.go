package controller

import (
	"strings"

	"github.com/Nexters/myply/domain/member"
	"github.com/gofiber/fiber/v2"
)

type MemberController interface {
	SignUp() fiber.Handler // TODO: pagination
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
				return c.Status(409).SendString("fail: account already exist")
			}

			return c.Status(500).SendString("fail: internal server error")
		}

		return c.Status(201).SendString("success: sign up")
	}
}
