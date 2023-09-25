package user

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	req := models.User{}
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	result, err := CreateUserService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func Login(c *gin.Context)  {
	req := models.User{}

	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	result, err := LoginService(req.Email, req.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", result, 60 * 60 * 2, "", "", false, true)

	c.IndentedJSON(http.StatusCreated, result)

}