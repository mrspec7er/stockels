package stock

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"stockels/app/object"
	"stockels/app/utils"
	"sync"
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



