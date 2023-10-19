package stock

import (
	"log"
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

func GetStocksReport(c *gin.Context){
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

	stocksBuffer, err := GetReportStockService(user, req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	

	_, err = c.Writer.Write(stocksBuffer.Bytes())
	if err != nil {
	log.Fatalln("Error in writing with context: ", err.Error())
	}
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