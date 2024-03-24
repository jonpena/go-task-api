package routes

import (
	"github.com/gofiber/fiber/v2"
	c "github.com/jonpena/api-task-go/controllers"
)

func UseRoutes(router fiber.Router) {
	route := router.Group("/tareas") 
	route.Get("", c.HandlerGetTasks)
	route.Get(":id", c.HandlerGetTask)
	route.Put(":id", c.HandlerUpdateTask)
	route.Post("", c.HandlerCreateTask)
	route.Delete(":id", c.HandlerDeleteTask)
}