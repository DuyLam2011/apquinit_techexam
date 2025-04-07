package main

import (
	"log"
	"os"

	"github.com/ether/v1/handlers"
	"github.com/ether/v1/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	utils.InitRedis()
	utils.InitMongo()

	app := fiber.New()

	// Enable CORS for all origins,
	// allowing all HTTP methods and headers(Change as needed)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))

	api := app.Group("/api/v1")
	api.Get("/eth/:address", handlers.GetEthInfo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
