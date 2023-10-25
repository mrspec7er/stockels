package subscribtion

import (
	"stockels/middleware"

	"github.com/gin-gonic/gin"
)

	func Routes(router *gin.Engine)  {
		routerGroup := router.Group("/api/v1")
		{
			routerGroup.POST("/subscribe", middleware.Authentication, SubscribeToStocks)
			routerGroup.GET("/subscribe", middleware.Authentication, GetSubscribtionStocks)
			routerGroup.POST("/subscribe/generate-report", middleware.Authentication, GenerateStockReport)
		}
	}