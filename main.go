package main

import (
	"log"

	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/database"
	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8080")
}
