package handlers

import (
	"log"
	"net/http"
	"stockels/app/object"
	"stockels/app/services/stock"

	"github.com/gin-gonic/gin"
)

func GetStocksReport(c *gin.Context){
	req := []*object.GetStockData{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	stocksBuffer, err := stock.GetReportStockService(req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	_, err = c.Writer.Write(stocksBuffer.Bytes())
	if err != nil {
	log.Fatalln("Error in writing with context: ", err.Error())
	}
}