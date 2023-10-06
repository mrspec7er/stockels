package user

import (
	"stockels/middleware"

	"github.com/gin-gonic/gin"
)

	func Routes(router *gin.Engine) {
		routerGroup := router.Group("/api/v1")
		{
			routerGroup.POST("/register", Register)
			routerGroup.POST("/login", Login)
			routerGroup.GET("/whoami", middleware.Authentication,whoami)
		}
	}