package stock

import (
	"net/http"
	"stockels/models"

	"github.com/gin-gonic/gin"
)

func AddNewStock(c *gin.Context){
	req := models.Stock{}
	err := c.Bind(&req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return
	}

	result := AddNewStockService(req.Symbol)

	c.IndentedJSON(http.StatusCreated, result)
}