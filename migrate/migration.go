package main

import (
	"fmt"
	"log"
	"stockels/app/models"
	"stockels/utils"

	"github.com/joho/godotenv"
)

func init()  {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	dbConnection := utils.DB()

	dbConnection.AutoMigrate(&models.User{}, &models.Stock{}, &models.Subscribtion{})
}

func main()  {
	fmt.Println("Successfully migrating model!")
}