package main

import (
	"github.com/emmanueluwa/gapi/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	//setting up routes
	setupRoutes(app)

	

	app.Listen(":3000")
}
