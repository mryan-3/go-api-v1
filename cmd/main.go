package main

import (
	"fmt"
	"github.com/go-api-v1/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main (){
    fmt.Println("Hello World Jones")

    app := fiber.New()
    database.ConnectDB()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello Jones ")
    })

    app.Listen(":3001")
}
