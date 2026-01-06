package middleware

import (
	"go-admin/internal/util"

	"github.com/gofiber/fiber/v3"
)

func CompanyIsolationMiddleware(c fiber.Ctx) error {
	companyID, err := GetCompanyID(c)
	if err != nil {
		return util.ForbiddenResponse(c, "Access denied")
	}
	if companyID.String() == "00000000-0000-0000-0000-000000000000" {
		return util.ForbiddenResponse(c, "Invalid company")
	}
	return c.Next()
}
