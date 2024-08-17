package handler

import (
	"github.com/evertonbzr/library_micro/internal/user/api/types"
	"github.com/evertonbzr/library_micro/internal/user/repository"
	"github.com/evertonbzr/library_micro/internal/user/util"
	"github.com/evertonbzr/library_micro/pkg/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		data := types.SignInRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := h.userRepository.GetUserByEmail(data.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email or password incorrect",
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email or password incorrect",
			})
		}

		token, _ := util.GenerateJwt(user)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
			"user":  user,
		})
	}
}

func (h *AuthHandler) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := types.SignUpRequest{}

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if data.FullName == "" || data.Email == "" || data.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "name, email, and password cannot be empty",
			})
		}

		userFound, _ := h.userRepository.GetUserByEmail(data.Email)
		if userFound.ID != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email already exists",
			})
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

		user := model.User{
			FullName: data.FullName,
			Email:    data.Email,
			Password: string(hashedPassword),
		}

		if err := h.userRepository.CreateUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		token, _ := util.GenerateJwt(&user)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
			"user":  user,
		})
	}
}
