package database

import (
	"fmt"
	"log"
	"instacart/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DNS = "host=ec2-107-22-238-112.compute-1.amazonaws.com user=bnsoflomemgdeq password=d334261218d6efb51f0a4681332284f99d5d809299576e4db6d388c11c65b1b6 dbname=de075m238kcaio port=5432"

 func Migration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("not connected to the database")
	}
	fmt.Print("connected to the database")
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Addtocart{})
}
