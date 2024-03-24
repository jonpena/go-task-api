package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jonpena/api-task-go/models"
	"github.com/jonpena/api-task-go/services"
)


func HandlerGetTasks(c *fiber.Ctx) error {
	return c.JSON(services.GetTasks())
}


func HandlerGetTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Empty String"})
	}
	return c.JSON(services.GetTask(id))
}


func HandlerCreateTask(c *fiber.Ctx) error {
	var body models.Tarea
	c.BodyParser(&body)
	return c.JSON(services.CreateTask(body))
}


func HandlerUpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Empty String"})
	}

	var body models.Tarea

	c.BodyParser(&body)

	updatedTasks := services.UpdateTask(id, body);

	return c.JSON(fiber.Map{"Updated Tasks": updatedTasks})
}


func HandlerDeleteTask(c *fiber.Ctx) error {

	id := c.Params("id")

  fmt.Println(id)

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Empty String"})
	}

	return c.JSON(fiber.Map{"Deleted Tasks": services.DeleteTask(id)})
}
