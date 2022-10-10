package main

import (
	"log"

	"assignment2/config"
	"assignment2/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    config.Connect()
    
    app.Get("/orders", handlers.GetOrders)
    app.Get("/orders/:Order_id", handlers.GetOrder)
    app.Post("/orders", handlers.CreateOrder)
    app.Put("/orders/:Order_id", handlers.UpdateOrder)
    app.Delete("/orders/:Order_id", handlers.DeleteOrder)

    log.Fatal(app.Listen(":3000"))
}