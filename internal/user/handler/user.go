package handler

import (
	"github.com/evertonbzr/library_micro/internal/user/repository"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRepository: repository.NewUserRepository(),
	}
}

func (h *UserHandler) GetMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Get Me")
	}
}
