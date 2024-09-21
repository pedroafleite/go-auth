package main

import (
	"log"

	"go-auth/database"
	"go-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
 
func main() {
    database.Connect()

    app := fiber.New()

    app.Use(cors.New(cors.Config{
      AllowCredentials: true,
      AllowOrigins:"http://localhost:3000",
    }))

    routes.Setup(app)

    log.Fatal(app.Listen(":3000"))
}
