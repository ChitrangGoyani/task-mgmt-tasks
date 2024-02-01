package routes

import (
	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Group("/api")
	app.Get("/tasks", controller.GetTasks)
	app.Post("/createTask", controller.CreateTask)
	app.Put("/updateTask/:id", controller.UpdateTask)
	app.Delete("/deleteTask/:id", controller.DeleteTask)
}
