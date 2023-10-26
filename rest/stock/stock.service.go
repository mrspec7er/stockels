package stock

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockels/models"
	"stockels/utils"
	"strconv"
	"time"
)
type StockDataType struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Sector string `json:"sector"`
	Logo string `json:"logo"`
	Website string `json:"website"`
}

type StockPriceType struct {
	Open string `json:"open"`
	High string `json:"high"`
	Low string `json:"low"`
	Close string `json:"close"`
	Volume string `json:"volume"`
	UpdatedAt string `json:"updated_at"`
}

type SubscribtionStockType struct {
	models.Stock
	models.Subscribtion
	SupportPercentage float32 `json:"supportPercentage"`
	ResistancePercentage float32 `json:"resistancePercentage"`
}

type StockDetailPriceType struct {
	Date string `json:"date"`
	Open string `json:"open"`
	High string `json:"high"`
	Low string `json:"low"`
	Close string `json:"close"`
	Volume int `json:"volume"`
}

type StockDetailType struct {
	Info models.Stock `json:"info"`
	Price []StockDetailPriceType `json:"price"`
}

type GoapiInformationResponseType struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data struct {
		Result *StockDataType `json:"result"`
		LastPrice *StockPriceType `json:"last_price"`
	} `json:"data"`
}

type GoapiPriceResponseType struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data struct {
		Count int `json:"count"`
		Results *[]StockDetailPriceType `json:"results"`
	} `json:"data"`
}

func GetMultipleStockService(subscribtions []models.Subscribtion) ([]SubscribtionStockType, error) {
	subStock := []SubscribtionStockType{}

	for _, sub := range subscribtions {

		stock, err := GetStockBySymbolService(sub.StockSymbol)
		if err != nil {
			break
		}

		closePrice, err := strconv.Atoi(stock.ClosePrice)
		if err != nil {
			break
		}
		subStock = append(subStock, SubscribtionStockType{Stock: stock, Subscribtion: sub, SupportPercentage: 100 - (float32(sub.SupportPrice) / float32(closePrice) * 100), ResistancePercentage: 100 - (float32(closePrice) / float32(sub.ResistancePrice) * 100)})

	}

	if len(subStock) == 0 {
		return subStock, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	return subStock, nil
}

func GetReportStockService(stocksReq []models.Subscribtion) (*bytes.Buffer, error) {
	subStock := []SubscribtionStockType{}

	for _, sub := range stocksReq {

		stock, err := GetStockBySymbolService(sub.StockSymbol)
		if err != nil {
			break
		}

		closePrice, err := strconv.Atoi(stock.ClosePrice)
		if err != nil {
			break
		}
		subStock = append(subStock, SubscribtionStockType{Stock: stock, Subscribtion: sub, SupportPercentage: 100 - (float32(sub.SupportPrice) / float32(closePrice) * 100), ResistancePercentage: 100 - (float32(closePrice) / float32(sub.ResistancePrice) * 100)})

	}

	if len(subStock) == 0 {
		return &bytes.Buffer{}, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	stocksRecords := [][]string{
		{"symbol", "name", "sector", "supportPercentage", "resistancePercentage", "supportPrice", "resistancePrice", "openPrice", "closePrice", "highestPrice", "lowestPrice", "volume", "lastUpdate", "website", "description"},
	}

	for _, record := range subStock {
		stocksRecords = append(stocksRecords, []string{record.Symbol, record.Name, record.Sector, PercentageFormat(record.SupportPercentage), PercentageFormat(record.ResistancePercentage), strconv.Itoa(record.SupportPrice), strconv.Itoa(record.ResistancePrice), record.OpenPrice, record.ClosePrice, record.HighestPrice, record.LowestPrice, record.Volume, record.LastUpdate, record.Website, record.Description})
	}

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(stocksRecords) 

	return csvBuffer, nil
}

func GetStockBySymbolService(symbol string) (models.Stock, error) {
	ctx := context.Background()
	stock := models.Stock{}

	cachedStock, err := utils.Cache().Get(ctx, symbol).Result()
	if err != nil {

		// Get data from goapi
		stock, err = GetStockInfoFromAPI(symbol)
		if err != nil {
			return models.Stock{}, err
		}

		err = CacheStock(symbol, stock)
		return stock, err
	}

	err = json.Unmarshal([]byte(cachedStock), &stock)

	return stock, err
}

func GetStockDetailService(symbol string, fromDate string, toDate string) (StockDetailType, error) {
	ctx := context.Background()
	stock := models.Stock{}
	stockPrice := []StockDetailPriceType{}
	stockDetail := StockDetailType{}

	cachedKey := symbol + "_" + fromDate + "_" + toDate

	cachedStockDetail, err := utils.Cache().Get(ctx, cachedKey).Result()
	if err != nil {

		// Get data from goapi
		stock, err = GetStockInfoFromAPI(symbol)
		if err != nil {
			return StockDetailType{}, err
		}

		stockPrice, err = GetStockPriceFromAPI(symbol, fromDate, toDate)
		if err != nil {
			return StockDetailType{}, err
		}

		stockDetail = StockDetailType{Info: stock, Price: stockPrice}

		err = CacheStockDetail(cachedKey, stockDetail)
		return stockDetail, err
	}

	err = json.Unmarshal([]byte(cachedStockDetail), &stockDetail)

	return stockDetail, err

}

func GetStockInfoFromAPI(symbol string) (models.Stock, error){
	fmt.Println("Fetching stock information with symbol: ", symbol, "to goapi.id")
	stockMetaData := GoapiInformationResponseType{}

	upstreamApiUrl := os.Getenv("UPSTREAM_API_URL")
	upstreamApiKey := os.Getenv("UPSTREAM_API_KEY")
	res, err := http.Get( upstreamApiUrl + symbol + upstreamApiKey)
	if err != nil {
		return models.Stock{}, err
	}

	stockStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Stock{}, err
	}

	err = json.Unmarshal(stockStreamMetaData, &stockMetaData)
	if err != nil || stockMetaData.Data.Result == nil {
		return models.Stock{}, errors.New("Failed to fetch data from goapi.id")
	}


	stock := models.Stock{
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

	err = utils.DB().Where(models.Stock{Symbol: symbol}).Assign(stock).FirstOrCreate(&stock).Error
	
	return stock, err
}

func GetStockPriceFromAPI(symbol string, fromDate string, toDate string) ([]StockDetailPriceType, error){
	fmt.Println("Fetching stock price with symbol: ", symbol, "to goapi.id")
	stockPriceMetaData := GoapiPriceResponseType{}

	upstreamApiUrl := os.Getenv("UPSTREAM_API_URL")
	upstreamApiKey := os.Getenv("UPSTREAM_API_KEY")

	res, err := http.Get(upstreamApiUrl + symbol + "/historical" + upstreamApiKey + "&from=" + fromDate + "&to=" + toDate)
	if err != nil {
		return []StockDetailPriceType{}, err
	}

	stockPriceStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return []StockDetailPriceType{}, err
	}

	err = json.Unmarshal(stockPriceStreamMetaData, &stockPriceMetaData)
	if err != nil || stockPriceMetaData.Data.Results == nil {
		return []StockDetailPriceType{}, errors.New("Failed to fetch price data from goapi.id")
	}

	stockPrice := stockPriceMetaData.Data.Results
	
	return *stockPrice, err
}

func CacheStock(key string, stock models.Stock) error {
	ctx := context.Background()
	stockStringified, err := json.Marshal(stock)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockStringified, time.Hour).Err()

	return err
}

func CacheStockDetail(key string, stockDetail StockDetailType) error {
	ctx := context.Background()
	stockDetailStringified, err := json.Marshal(stockDetail)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockDetailStringified, time.Hour).Err()

	return err
}

func PercentageFormat(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 64)
}