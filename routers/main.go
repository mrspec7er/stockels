package routers

import (
	"stockels/module/stock"
	"stockels/module/user"

	"github.com/gin-gonic/gin"
)

func Config()  {
	router := gin.Default()
	stock.Routes(router)
	user.Routes(router)
	router.Run()
}