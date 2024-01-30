package main

import (
	"log"
	"os"

	"github.com/gabriel-kimutai/ano2tes/controller"
	"github.com/gabriel-kimutai/ano2tes/database"
	"github.com/gabriel-kimutai/ano2tes/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDB()
	file, err := os.OpenFile("out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	logger := logger.New(
		logger.Config{
			Output: file,
		},
	)
	app.Static("/static", "./public")
	app.Use(logger)

	app.Get("/", controller.Root)
	// User route
	app.Get("/user", middleware.RequireAuth, controller.User)
	app.Post("/user", controller.UserCreate)

	app.Get("/login", controller.LoginPage)
	app.Post("/login", controller.UserLogin)

	app.Get("/logout", controller.UserLogout)

	app.Listen(":3000")
}
