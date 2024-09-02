package main

import (
	"log"

	"go-auth/database"
	"go-auth/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
  }
 
func main() {
    database.Connect()

    app := fiber.New()

    routes.Setup(app)

    log.Fatal(app.Listen(":3000"))
}
