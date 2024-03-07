package handler

import (
	"auth/internal/core/domain"
	"auth/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(AuthService services.AuthService) *AuthHandler {
	return &AuthHandler{service: AuthService}
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	var user domain.Users
	if err := c.BodyParser(&user); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	if _, err := h.service.SignUp(&user); err != nil {
		return HandlerError(c, fiber.StatusBadRequest, err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
