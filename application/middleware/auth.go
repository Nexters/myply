package middleware

import (
	"strings"

	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/domain/member"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	New() fiber.Handler
}

type authMiddleware struct {
	memberRepository member.MemberRepository
	Targets          []string
}

func (a *authMiddleware) New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !a.isTarget(c.Method(), c.Path()) {
			return c.Next()
		}

		deviceToken := c.Get("Device-Token")
		member, err := a.memberRepository.Get(deviceToken)
		if err != nil {
			return c.Status(401).JSON(controller.BaseResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		c.Locals("member", member)

		return c.Next()
	}
}

func (a *authMiddleware) isTarget(method, path string) bool {
	targets := a.Targets
	target := strings.ToUpper(method) + " " + path

	for i := range targets {
		if targets[i] == target {
			return true
		}
	}

	return false
}

func NewAuthMiddleware(memberRepository member.MemberRepository) AuthMiddleware {
	return &authMiddleware{
		memberRepository: memberRepository,
		Targets: []string{
			"GET /api/v1/members/",
		},
	}
}
