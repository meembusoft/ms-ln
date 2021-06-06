package handler

import (
	"ashikCodecamper/lagbenaki/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": nil})
}

// GetProduct query product
func GetProduct(c *fiber.Ctx) error {

	if "" == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": nil})
}

// CreateProduct new product
func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	fmt.Println(product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) error {

	if "" == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
