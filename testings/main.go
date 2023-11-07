package testings

import (
	"log"
	"stockels/app/routers"
	"stockels/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func SetupRouters(configFilePath string) *gin.Engine {
	err := godotenv.Load(configFilePath)

	if err != nil {
		log.Fatal(err)
	}
	
	utils.DB()
	utils.Cache()
	return routers.Config()
}