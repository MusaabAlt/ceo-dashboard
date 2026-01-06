package router

import (
	"go-admin/internal/config"
	"go-admin/internal/handler"
	"go-admin/internal/middleware"
	"go-admin/internal/service"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	db := config.GetDB()
	authService := service.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	protected := api.Group("", middleware.AuthMiddleware)
	protected.Use(middleware.CompanyIsolationMiddleware)
}
