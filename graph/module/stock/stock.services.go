package stock

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"stockels/graph/object"
	"stockels/models"
	"stockels/utils"
	"strconv"
	"sync"
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
		Results []*object.StockDetailPrice `json:"results"`
	} `json:"data"`
}

func GetMultipleStockService(stocksReq []*object.GetStockData) ([]*object.StockData, error) {
	stocks := []*object.StockData{}

	stockCtx := make(chan *object.StockData, len(stocksReq))
	wg := &sync.WaitGroup{}
	wg.Add(len(stocksReq))

	for _, sub := range stocksReq {
		go AsyncGetStockService(sub.StockSymbol, sub.SupportPrice, sub.ResistancePrice, stockCtx, wg)
	}

	wg.Wait()
	close(stockCtx)

	for stock := range stockCtx {
		stocks = append(stocks, &object.StockData{Name: stock.Name, Symbol: stock.Symbol, Description: stock.Description, Sector: stock.Sector, Logo: stock.Logo, Website: stock.Website, OpenPrice: stock.OpenPrice, ClosePrice: stock.ClosePrice, HighestPrice: stock.HighestPrice, LowestPrice: stock.LowestPrice, Volume: stock.Volume, LastUpdate: stock.LastUpdate, SupportPercentage: stock.SupportPercentage, ResistancePercentage: stock.ResistancePercentage})
	}

	if len(stocks) == 0 {
		return stocks, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	return stocks, nil
}


func GetStockBySymbolService( symbol string, supportPrice int, resistancePrice int) (*object.StockData, error) {
	ctx := context.Background()
	stock := &object.StockData{}

	cachedStock, err := utils.Cache().Get(ctx, symbol).Result()
	if err != nil {

		// Get data from goapi
		stock, err = GetStockInfoFromAPI(symbol)
		if err != nil {
			return stock, err
		}

		supportPercentage, resistancePercentage, err := GetSupportAndResistancePercentage(stock.ClosePrice, supportPrice, resistancePrice)
		if err != nil {
			return &object.StockData{}, err
		}

		stock.SupportPercentage = *supportPercentage
		stock.ResistancePercentage = *resistancePercentage

		err = CacheStock(symbol, stock)
		return stock, err
	}

	err = json.Unmarshal([]byte(cachedStock), &stock)

	supportPercentage, resistancePercentage, err := GetSupportAndResistancePercentage(stock.ClosePrice, supportPrice, resistancePrice)
	if err != nil {
		return &object.StockData{}, err
	}

	stock.SupportPercentage = *supportPercentage
	stock.ResistancePercentage = *resistancePercentage

	return stock, err
}

func GetStockDetailService(symbol string, fromDate string, toDate string, supportPrice int, resistancePrice int) (*object.StockDetail, error) {
	ctx := context.Background()
	stock := &object.StockData{}
	stockPrice := []*object.StockDetailPrice{}
	stockDetail := &object.StockDetail{}

	cachedKey := symbol + "_" + fromDate + "_" + toDate

	cachedStockDetail, err := utils.Cache().Get(ctx, cachedKey).Result()
	if err != nil {

		// Get data from goapi
		stock, err = GetStockInfoFromAPI(symbol)
		if err != nil {
			return &object.StockDetail{}, err
		}

		supportPercentage, resistancePercentage, err := GetSupportAndResistancePercentage(stock.ClosePrice, supportPrice, resistancePrice)
		if err != nil {
			return &object.StockDetail{}, err
		}

		stock.SupportPercentage = *supportPercentage
		stock.ResistancePercentage = *resistancePercentage

		stockPrice, err = GetStockPriceFromAPI(symbol, fromDate, toDate)
		if err != nil {
			return &object.StockDetail{}, err
		}

		stockDetail = &object.StockDetail{Info: stock, Price: stockPrice}

		err = CacheStockDetail(cachedKey, stockDetail)
		return stockDetail, err
	}

	err = json.Unmarshal([]byte(cachedStockDetail), &stockDetail)

	return stockDetail, err

}

func GetReportStockService(stocksReq []*object.GetStockData) (*bytes.Buffer, error) {
	
	subStock, err := GetMultipleStockService(stocksReq)

	if err != nil {
		return &bytes.Buffer{}, errors.New("Failed to get data from 'GetMultipleStockService'!")
	}

	stocksRecords := [][]string{
		{"symbol", "name", "sector", "supportPercentage", "resistancePercentage", "openPrice", "closePrice", "highestPrice", "lowestPrice", "volume", "lastUpdate", "website", "description"},
	}

	for _, record := range subStock {
		stocksRecords = append(stocksRecords, []string{record.Symbol, record.Name, record.Sector, PercentageFormat(float32(record.SupportPercentage)), PercentageFormat(float32(record.ResistancePercentage)), record.OpenPrice, record.ClosePrice, record.HighestPrice, record.LowestPrice, record.Volume, record.LastUpdate, record.Website, record.Description})
	}

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(stocksRecords) 

	return csvBuffer, nil
}

func AsyncGetStockService( symbol string, supportPrice int, resistancePrice int, stockCtx chan *object.StockData, wg *sync.WaitGroup) {
	ctx := context.Background()
	stock := &object.StockData{}

	cachedStock, err := utils.Cache().Get(ctx, symbol).Result()
	if err != nil {

		// Get data from goapi
		stock, err = GetStockInfoFromAPI(symbol)
		if err != nil {
			log.Println(err.Error())
			return
		}

		supportPercentage, resistancePercentage, err := GetSupportAndResistancePercentage(stock.ClosePrice, supportPrice, resistancePrice)
		if err != nil {
			log.Println(err.Error())
			return
		}

		stock.SupportPercentage = *supportPercentage
		stock.ResistancePercentage = *resistancePercentage

		err = CacheStock(symbol, stock)
		if err != nil {
			log.Println(err.Error())
			return
		}

		stockCtx <- stock
		wg.Done()
		return
	}

	err = json.Unmarshal([]byte(cachedStock), &stock)

	supportPercentage, resistancePercentage, err := GetSupportAndResistancePercentage(stock.ClosePrice, supportPrice, resistancePrice)
	if err != nil {
		log.Println(err.Error())
			return
	}

	stock.SupportPercentage = *supportPercentage
	stock.ResistancePercentage = *resistancePercentage

	stockCtx <- stock
	wg.Done()
	return 
}

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

	err = utils.DB().Where(models.Stock{Symbol: symbol}).Assign(stock).FirstOrCreate(&stock).Error
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



