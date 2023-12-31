package handlers

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine)  {
	routerGroup := router.Group("/api/v1")
	{
		routerGroup.POST("/stocks/generate-report", GetStocksReport)
	}
}