package main

import (
	"log"
	"stockels/routers"
	"stockels/utils"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	utils.DB()
	routers.Config()
}