package models

type Product struct {
	ProductID  	uint   `json:"product_id" gorm:"primaryKey; autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Images		string `json:"image"`
	Price       int    `json:"price"`
	Stars       string `json:"stars"`
	Quantity    int    `json:"quantity"`
}

