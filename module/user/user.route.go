package user

import "github.com/gin-gonic/gin"

	func Routes(router *gin.Engine) {
		routerGroup := router.Group("/api/v1")
		{
			routerGroup.POST("/users", Register)
		}
	}