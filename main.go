package main

import (
	"fmt"
	"instacart/database"
	"instacart/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}

func Routes(app *fiber.App) {

	//customerregister
	app.Post("/user", route.CustomerReg)
	app.Get("/user", route.GetCustomer)
	//Product
	app.Post("/product", route.AddProduct)
	app.Get("/product", route.GetProductName)
	app.Get("/product/:id", route.GetProduct)
	app.Delete("/product/:id", route.Delete)
	app.Put("/product/:id", route.Update)
	//addtocart
	app.Post("/addtocart", route.AddToCart)
	app.Get("/getcart/:userID", route.GetCart)
	//logincustomer
	app.Post("/login",route.Log)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Could not load .env file")
	}
	PORT := os.Getenv("PORT")
	database.Migration()
	app := fiber.New()
	Routes(app)
	log.Fatal(app.Listen(":" + PORT))
}