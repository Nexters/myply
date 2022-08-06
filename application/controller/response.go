package controller

import "github.com/gofiber/fiber/v2"

type Response struct {
	code    int32
	message string
	data    fiber.Map
}

func (r *Response) toMap() fiber.Map {
	return fiber.Map{
		"code":    r.code,
		"message": r.message,
		"data":    r.data,
	}
}
