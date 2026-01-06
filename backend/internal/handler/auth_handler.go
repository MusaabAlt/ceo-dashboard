package handler

import (
	"go-admin/internal/dto"
	"go-admin/internal/service"
	"go-admin/internal/util"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.Bind().JSON(&req); err != nil {
		return util.ValidationErrorResponse(c, "Invalid body")
	}
	response, err := h.authService.Register(&req)
	if err != nil {
		return util.ErrorResponse(c, 400, err.Error())
	}
	return util.SuccessResponse(c, response, "Registered")
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.Bind().JSON(&req); err != nil {
		return util.ValidationErrorResponse(c, "Invalid body")
	}
	response, err := h.authService.Login(&req)
	if err != nil {
		return util.ErrorResponse(c, 401, err.Error())
	}
	return util.SuccessResponse(c, response, "Logged in")
}
