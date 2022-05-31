package models

type Carts struct {
	Products map[string]interface{} `json:"products"`
	MyUser   User                   `json:"user"`
}