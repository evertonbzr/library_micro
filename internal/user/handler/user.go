package handler

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Get Me")
	}
}
