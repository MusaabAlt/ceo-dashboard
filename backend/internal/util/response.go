package util

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c fiber.Ctx, data interface{}, message string) error {
	return c.Status(200).JSON(Response{Success: true, Message: message, Data: data})
}

func ErrorResponse(c fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{Success: false, Error: message})
}

func ValidationErrorResponse(c fiber.Ctx, message string) error {
	return ErrorResponse(c, 400, message)
}

func UnauthorizedResponse(c fiber.Ctx, message string) error {
	return ErrorResponse(c, 401, message)
}

func ForbiddenResponse(c fiber.Ctx, message string) error {
	return ErrorResponse(c, 403, message)
}

func NotFoundResponse(c fiber.Ctx, message string) error {
	return ErrorResponse(c, 404, message)
}

func InternalErrorResponse(c fiber.Ctx, message string) error {
	return ErrorResponse(c, 500, message)
}
