package main

import (
	"log"
	"stockels/app/routers"
	"stockels/app/utils"

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
	utils.Cache()
	app := routers.Config()

	app.Run()
}