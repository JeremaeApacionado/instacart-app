package models

type Addtocart struct {
	ID              uint `json:"id" gorm:"foreignKey"`
	UserID          uint `json:"userID"`
	ProductID       uint `json:"productID"`
	QuantityOrdered uint `json:"quantityOrdered"`
}