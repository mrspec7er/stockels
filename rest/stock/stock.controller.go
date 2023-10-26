package stock

import (
	"log"
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

	result, err := GetMultipleStockService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	c.IndentedJSON(http.StatusCreated, result)

}

func GetStocksReport(c *gin.Context){
	req := []models.Subscribtion{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	stocksBuffer, err := GetReportStockService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	_, err = c.Writer.Write(stocksBuffer.Bytes())
	if err != nil {
	log.Fatalln("Error in writing with context: ", err.Error())
	}
}

func GetStockDetail(c *gin.Context){
	symbol := c.Param("symbol");
	fromDate := c.Query("from");
	toDate := c.Query("to")

	result, err := GetStockDetailService(symbol, fromDate, toDate)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}