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

    // // Migrate the schema
    // db.AutoMigrate(&Product{})
    // db.Create(&Product{Code: "D42", Price: 100})

    app := fiber.New()

    routes.Setup(app)


    log.Fatal(app.Listen(":3000"))
}
