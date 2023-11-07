package stock_test

import (
	"reflect"
	"stockels/app/object"
	"stockels/app/services/stock"
	"stockels/testings"
	"testing"
)

func init()  {
	testings.SetupRouters("../../../.env")
}

func TestStockServices(t *testing.T)  {
	t.Run("should return stock info", func(t *testing.T) {
		stocks := &object.StockData{}

		result, err := stock.GetStockInfoFromAPI("BBCA")
	
		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stocks)) {
			t.Errorf(err.Error())
		}
	})

	t.Run("should return stock price", func(t *testing.T) {
		stockDetail := []*object.StockDetailPrice{}

		result, err := stock.GetStockPriceFromAPI("BBNI", "2023-05-11", "2023-11-2")

		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stockDetail)) {
			t.Errorf(err.Error())
		}
	})

	t.Run("should return array of stock info", func(t *testing.T) {
		getStocksPayload := []*object.GetStockData{{StockSymbol: "INDF", SupportPrice: 6275, ResistancePrice: 7550}, {StockSymbol: "ASII", SupportPrice: 5380, ResistancePrice: 6750}}

		stocks := []*object.StockData{}

		result, err := stock.GetMultipleStockService(getStocksPayload)

		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stocks)) {
			t.Errorf(err.Error())
		}
	})
}