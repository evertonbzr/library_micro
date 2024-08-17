package handler

import (
	"github.com/evertonbzr/library_micro/internal/user/repository"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userRepository repository.UserRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userRepository: *repository.NewUserRepository(),
	}
}

func (h *AuthHandler) SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Sign In")
	}
}

func (h *AuthHandler) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Sign Up")
	}
}
