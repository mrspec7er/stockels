package routers

import (
	"stockels/module/stock"

	"github.com/gin-gonic/gin"
)

func Config()  {
	router := gin.Default()
	stock.Routes(router)
	router.Run()
}