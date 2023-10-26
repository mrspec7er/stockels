package subscribtion

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func SubscribeToStocks(c *gin.Context){
	req := []models.Subscribtion{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	userCtx, status := c.Get("user")
	
	if !status  {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to Get User Data!"});
		return
	}

	user := userCtx.(models.User)

	result, err := SubscribeMultipleStockService(req, user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	c.IndentedJSON(http.StatusCreated, result)

}

func GetSubscribtionStocks(c *gin.Context){
	req := []models.Subscribtion{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	userCtx, status := c.Get("user")
	
	if !status  {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to Get User Data!"});
		return
	}

	user := userCtx.(models.User)

	result, err := GetSubscribtionStockService(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	c.IndentedJSON(http.StatusCreated, result)

}

func GenerateStockReport(c *gin.Context){
	userCtx, status := c.Get("user")
	
	if !status  {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to Get User Data!"});
		return
	}

	user := userCtx.(models.User)
	reportUrl, err := GenerateStockReportService(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	c.IndentedJSON(http.StatusCreated, reportUrl)
}