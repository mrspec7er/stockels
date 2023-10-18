package stock

import (
	"stockels/middleware"

	"github.com/gin-gonic/gin"
)

	func Routes(router *gin.Engine)  {
		routerGroup := router.Group("/api/v1")
		{
			routerGroup.POST("/stocks", GetStocks)
			routerGroup.GET("/stocks/:symbol", GetStockDetail)
			routerGroup.POST("/stocks/subscribe", middleware.Authentication, SubscribeToStocks)
			routerGroup.GET("/stocks/subscribe", middleware.Authentication, GetSubscribtionStocks)
			routerGroup.GET("/stocks/report", middleware.Authentication, GetSubscribtionStocksReport)
			routerGroup.POST("/stocks/generate-report", middleware.Authentication, GenerateStockReport)
		}
	}