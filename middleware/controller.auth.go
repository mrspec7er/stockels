package middleware

import (
	"fmt"
	"net/http"
	"os"
	"stockels/models"
	"stockels/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context)  {
	token, err := c.Cookie("Authorization");

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	claims, ok := payload.Claims.(jwt.MapClaims)
	if ok && payload.Valid {
		user := models.User{}
		err := utils.DB().Find(&user, claims["id"]).Error
		if err != nil || user.ID == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Cannot Find User Object!"});
			return
		}
		c.Set("user", user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}