package stock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockels/app/models"
	"stockels/app/object"
	"stockels/app/utils"
	"strconv"
	"time"
)

func GetStockInfoFromAPI(symbol string) (*object.StockData, error){
	fmt.Println("Fetching stock information with symbol: ", symbol, "to goapi.id")
	stockMetaData := GoapiInformationResponseType{}

	upstreamApiUrl := os.Getenv("UPSTREAM_API_URL")
	upstreamApiKey := os.Getenv("UPSTREAM_API_KEY")
	res, err := http.Get( upstreamApiUrl + symbol + upstreamApiKey)
	if err != nil {
		return &object.StockData{}, err
	}

	stockStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return &object.StockData{}, err
	}

	err = json.Unmarshal(stockStreamMetaData, &stockMetaData)
	if err != nil || stockMetaData.Data.Result == nil {
		return &object.StockData{}, errors.New("Failed to fetch data from goapi.id")
	}


	stock := &models.Stock{
		Symbol: symbol,
		Name: stockMetaData.Data.Result.Name,
		Sector: stockMetaData.Data.Result.Sector,
		Logo: stockMetaData.Data.Result.Logo,
		Website: stockMetaData.Data.Result.Website,
		Description: stockMetaData.Data.Result.Description,
		OpenPrice: stockMetaData.Data.LastPrice.Open,
		ClosePrice: stockMetaData.Data.LastPrice.Close,
		HighestPrice: stockMetaData.Data.LastPrice.High,
		LowestPrice: stockMetaData.Data.LastPrice.Low,
		Volume: stockMetaData.Data.LastPrice.Volume,
		LastUpdate: stockMetaData.Data.LastPrice.UpdatedAt,
	}

	err = utils.DB().Where(models.Stock{Symbol: symbol}).Assign(stock).Save(&stock).Error
	result := &object.StockData{Name: stock.Name, Symbol: stock.Symbol, Description: stock.Description, Sector: stock.Sector, Logo: stock.Logo, Website: stock.Website, OpenPrice: stock.OpenPrice, ClosePrice: stock.ClosePrice, HighestPrice: stock.HighestPrice, LowestPrice: stock.LowestPrice, Volume: stock.Volume, LastUpdate: stock.LastUpdate}
	
	return result, err
}

func GetStockPriceFromAPI(symbol string, fromDate string, toDate string) ([]*object.StockDetailPrice, error){
	fmt.Println("Fetching stock price with symbol: ", symbol, "to goapi.id")
	stockPriceMetaData := GoapiPriceResponseType{}

	upstreamApiUrl := os.Getenv("UPSTREAM_API_URL")
	upstreamApiKey := os.Getenv("UPSTREAM_API_KEY")

	res, err := http.Get(upstreamApiUrl + symbol + "/historical" + upstreamApiKey + "&from=" + fromDate + "&to=" + toDate)
	if err != nil {
		return []*object.StockDetailPrice{}, err
	}

	stockPriceStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return []*object.StockDetailPrice{}, err
	}

	err = json.Unmarshal(stockPriceStreamMetaData, &stockPriceMetaData)
	if err != nil || stockPriceMetaData.Data.Results == nil {
		return []*object.StockDetailPrice{}, errors.New("Failed to fetch price data from goapi.id")
	}

	stockPrice := stockPriceMetaData.Data.Results
	
	return stockPrice, err
}

func CacheStock(key string, stock *object.StockData) error {
	ctx := context.Background()
	stockStringified, err := json.Marshal(stock)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockStringified, time.Hour).Err()

	return err
}

func CacheStockDetail(key string, stockDetail *object.StockDetail) error {
	ctx := context.Background()
	stockDetailStringified, err := json.Marshal(stockDetail)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockDetailStringified, time.Hour).Err()

	return err
}

func GetSupportAndResistancePercentage(closePriceStr string, supportPrice int, resistancePrice int) (*float64, *float64, error)  {
	closePrice, err := strconv.Atoi(closePriceStr)
	if err != nil {
		return nil, nil, err
	}

	supportPercentage := 100 - (float64(supportPrice) / float64(closePrice) * 100)
	resistancePercentage := 100 - (float64(closePrice) / float64(resistancePrice) * 100)

	return &supportPercentage, &resistancePercentage, nil
}

func PercentageFormat(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 64)
}