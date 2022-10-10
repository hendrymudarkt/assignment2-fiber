package handlers

import (
	"assignment2/config"
	"assignment2/entities"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
    var orders []entities.Orders

    config.Database.Preload("Items").Find(&orders)
    return c.Status(200).JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
    id := c.Params("order_id")
    var order entities.Orders

    result := config.Database.Preload("Items").Find(&order, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&order)
}

func CreateOrder(c *fiber.Ctx) error {
    order := new(entities.Orders)

    if err := c.BodyParser(order); err != nil {
        return c.Status(503).SendString(err.Error())
    }
	
    config.Database.Create(&order)
    return c.Status(201).JSON(order)
}

func UpdateOrder(c *fiber.Ctx) error {
    order := new(entities.Orders)
    id := c.Params("order_id")

    if err := c.BodyParser(order); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    sqlQuery := `
    UPDATE orders INNER JOIN items ON orders.order_id = items.order_id
    SET customer_name = ?, ordered_at = ?, quantity = ?
    WHERE orders.order_id = ?
    `

    config.Database.Exec(sqlQuery, order.Customer_name, order.Ordered_at, order.Items[0].Quantity, id)

    return c.Status(200).JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
    id := c.Params("order_id")
    var order entities.Orders

    result := config.Database.Delete(&order, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}