package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/suraj1294/go-fiber-planetscale/handler"
)

type Address struct {
	Street string
	City   string
	Pin    int
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	fmt.Println("hello")

	app := fiber.New()

	app.Use(logger.New())

	ph := handler.GetProductHandler()

	app.Get("/products", ph.GetProducts)
	app.Get("/products/:id", ph.GetProduct)
	app.Post("/products", ph.AddProduct)
	app.Put("/products/:id", ph.UpdateProduct)
	app.Delete("/products/:id", ph.DeleteProduct)

	app.Listen(":8080")
}
