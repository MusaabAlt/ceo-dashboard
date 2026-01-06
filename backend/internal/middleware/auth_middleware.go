package middleware

import (
	"go-admin/internal/config"
	"go-admin/internal/util"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func AuthMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return util.UnauthorizedResponse(c, "Missing auth header")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := util.ValidateJWT(tokenString, config.AppConfig.JWTSecret)
	if err != nil {
		return util.UnauthorizedResponse(c, "Invalid token")
	}
	c.Locals("user_id", claims.UserID)
	c.Locals("company_id", claims.CompanyID)
	c.Locals("email", claims.Email)
	return c.Next()
}

func GetUserID(c fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(401, "Not authenticated")
	}
	return userID, nil
}

func GetCompanyID(c fiber.Ctx) (uuid.UUID, error) {
	companyID, ok := c.Locals("company_id").(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(401, "No company")
	}
	return companyID, nil
}
