package stock

import (
	"fmt"
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

func GetStockDetail(c *gin.Context){
	symbol := c.Param("symbol");
	fromDate := c.Query("from");
	toDate := c.Query("to")

	fmt.Println(fromDate, toDate)

	result, err := GetStockDetailService(symbol, fromDate, toDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}