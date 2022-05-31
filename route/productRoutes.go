package route

import (
	"errors"
	"instacart/database"
	"instacart/models"
	"github.com/gofiber/fiber/v2"
)

//foodlist
func AddProduct(c *fiber.Ctx) error {
	var food models.Product
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}

	database.DB.Create(&food)
	return c.JSON(&food)
}

func GetProductName(c *fiber.Ctx) error {
	var food []models.Product

	database.DB.Find(&food)
	return c.JSON(&food)
}

func FindProduct(id int, food *models.Product) error {
	database.DB.Find(&food, "id=?", id)
	if food.ProductID == 0 {
		return errors.New("ProductId does not existed")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var food models.Product
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(&food)

}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food models.Product
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Delete(&food)
	return c.SendString("Deleted product")
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food models.Product
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Save(&food)
	return c.JSON(&food)
}