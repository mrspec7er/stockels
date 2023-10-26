package controller

import (
	"log"
	"net/http"
	"stockels/graph/module/stock"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func GetStocksReport(c *gin.Context){
	req := []*models.Subscribtion{}
	
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

func Routes(router *gin.Engine)  {
	routerGroup := router.Group("/api/v1")
	{
		routerGroup.POST("/stocks/generate-report", GetStocksReport)
	}
}