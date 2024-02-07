package routes

import (
	"github.com/ChitrangGoyani/task-mgmt-tasks-backend.git/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/tasks", controller.GetTasks)
	api.Post("/createTask", controller.CreateTask)
	api.Put("/updateTask/:id", controller.UpdateTask)
	api.Delete("/deleteTask/:id", controller.DeleteTask)
	api.Get("/searchTask/:content", controller.SearchTask)
}
