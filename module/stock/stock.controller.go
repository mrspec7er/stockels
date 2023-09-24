package stock

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func GetStocks(c *gin.Context){
	req := []models.Subscribtion{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	result, err := GetAllStockServices(req)
	c.IndentedJSON(http.StatusCreated, result)

}

func GetStockDetail(c *gin.Context){
	symbol := c.Param("symbol")

	result, err := GetEachStockServices(symbol)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}