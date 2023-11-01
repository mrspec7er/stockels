package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"stockels/app/models"
	"stockels/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if values, _ := c.Request.Header["Authorization"]; len(values) == 0 {
			c.Next()
			return
		}
		token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]


		payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		claims, ok := payload.Claims.(jwt.MapClaims)

		if ok && payload.Valid {
			user := models.User{}
			err := utils.DB().Find(&user, claims["id"]).Error
			if err != nil || user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized);
				return
			}
			ctx := context.WithValue(c.Request.Context(), "user", user)
			c.Request = c.Request.WithContext(ctx)
			
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
	}
}

func GetAuthContextMiddleware(ctx context.Context) (*models.User, error) {
	ginContext := ctx.Value("user")
	if ginContext == nil {
		err := fmt.Errorf("Unauthorize user")
		return nil, err
	}

	gc, ok := ginContext.(models.User)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return &gc, nil
}