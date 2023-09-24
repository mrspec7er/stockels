package stock

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type StockReqType struct {
	StockList []string `json:"stockList"`
}

func GetStocks(c *gin.Context){
	req := StockReqType{}
	
	err := c.Bind(&req)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}
	fmt.Println("REQ: ", req.StockList)

	result, err := GetAllStockServices(req.StockList)
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