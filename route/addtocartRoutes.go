package route

import (
	"instacart/database"
	"instacart/models"

	"github.com/gofiber/fiber/v2"
)

//add to cart
func AddToCart(c *fiber.Ctx) error {
	cart := models.Addtocart{}
	if err := c.BodyParser(&cart); err != nil {
		return err
	}
	result := database.DB.Debug().Create(&cart)
	if result.RowsAffected > 0 {
		return c.JSON("Success insert")
	}

	return c.JSON(result.Error)
	
}

func getUserInfo(userID int) map[string]interface{} {
	user := make(map[string]interface{})
	database.DB.Table("users").Select("user_id, fullname, email, address").Find(&user, "user_id", userID)
	return user
}


func getProductDetails(product_id []uint) []map[string]interface{} {
	// product := make(map[string]interface{})
	result := []map[string]interface{}{}
	// product := []models.Product{}
	database.DB.Debug().Table("products p").
		Select("p.product_id, p.name, p.description, p.images, p.stars, c.quantity_ordered").
		Joins("JOIN addtocarts c ON p.product_id = c.product_id").Where("p.product_id IN ?", product_id).Find(&result)
	// fmt.Println(product)
	return result
}


func GetCart(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("userID")
	je := []models.Addtocart{}
	user := make(map[string]interface{})
	cartProduct := []map[string]interface{}{}
	database.DB.Debug().Find(&je, "user_id = ?", userID)
	mapInterface := make(map[string]interface{})
	if len(je) > 0 {
		user = getUserInfo(userID)
		prodIDs := []uint{}
		for _, element := range je {
			prodIDs = append(prodIDs, element.ProductID)
		}
		cartProduct = getProductDetails(prodIDs)

		// var b = models.Carts{Products: cartProduct, MyUser: user}
		// fmt.Println(b)
		mapInterface["user"] = user
		mapInterface["cart"] = cartProduct
		data := models.Result{Code: "00", Message: "Success response", Data: mapInterface}
		return c.JSON(data)
	}
	data := models.Result{Code: "01", Message: "No cart", Data: mapInterface}
	return c.JSON(data)
}