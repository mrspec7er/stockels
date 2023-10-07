package stock

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

type SubscribtionStockType struct {
	models.Stock
	models.Subscribtion
	SupportPercentage float32 `json:"supportPercentage"`
	ResistancePercentage float32 `json:"resistancePercentage"`
}

func GetStocks(c *gin.Context){
	req := []models.Subscribtion{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	result, err := GetMultipleStockService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	c.IndentedJSON(http.StatusCreated, result)

}

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

func GetStockDetail(c *gin.Context){
	symbol := c.Param("symbol")

	result, err := GetStockBySymbolService(symbol)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}