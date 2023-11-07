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

func TestGetStockFromAPI(t *testing.T)  {
	stocks := &object.StockData{}

	result, err := stock.GetStockInfoFromAPI("BBCA")

	if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stocks)) {
		t.Errorf(err.Error())
	}

}

func TestGetMultipleStock(t *testing.T)  {
	getStocksPayload := []*object.GetStockData{{StockSymbol: "INDF", SupportPrice: 6275, ResistancePrice: 7550}, {StockSymbol: "ASII", SupportPrice: 5380, ResistancePrice: 6750}}

	stocks := []*object.StockData{}

	result, err := stock.GetMultipleStockService(getStocksPayload)

	if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stocks)) {
		t.Errorf(err.Error())
	}

}