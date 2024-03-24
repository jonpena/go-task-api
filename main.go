package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/jonpena/api-task-go/config"
	"github.com/jonpena/api-task-go/routes"
)


func main() {
	
	godotenv.Load()
	
	app := fiber.New()

	app.Use(cors.New())
	
	app.Use(logger.New());
	
	config.ConnectToMongo()
	
	routes.UseRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}