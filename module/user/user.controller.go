package user

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	req := models.User{}
	c.Bind(&req)

	result, err := CreateUserService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}