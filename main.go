package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

var defaultPort = "8080"

func main() {

	mode := os.Getenv("GIN_MODE")

	if mode != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load env", err)
		}
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	app.Use(logger.New())

	ph := handler.GetProductHandler()

	app.Get("/products", ph.GetProducts)
	app.Get("/products/:id", ph.GetProduct)
	app.Post("/products", ph.AddProduct)
	app.Put("/products/:id", ph.UpdateProduct)
	app.Delete("/products/:id", ph.DeleteProduct)

	app.Listen(fmt.Sprintf(":%s", port))
}
