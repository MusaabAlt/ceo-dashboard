package main

import (
	"log"

	"go-admin/internal/config"
	"go-admin/internal/domain"
	"go-admin/internal/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()
	migrateDatabase()

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
		AllowMethods: []string{"GET, POST, PUT, DELETE, PATCH"},
	}))

	router.SetupRoutes(app)

	port := ":" + config.AppConfig.Port
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(app.Listen(port))
}

func customErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

func migrateDatabase() {
	db := config.GetDB()
	err := db.AutoMigrate(
		&domain.Company{},
		&domain.Role{},
		&domain.User{},
		&domain.UserRole{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("✅ Database migration completed")
}
