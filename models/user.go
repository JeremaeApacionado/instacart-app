package models

type User struct {
	UserID       	uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Fullname 	  	string `json:"Fullname"`
	Contact			string `json:"contact_no"`
	Email   	  	string `json:"Email"`
	Address		  	string `json:"Address"`
	Username	  	string `json:"Username"`
	Password  	  	string `json:"Password"`
}
