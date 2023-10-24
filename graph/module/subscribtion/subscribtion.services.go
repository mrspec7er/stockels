package subscribtion

import (
	"bytes"
	"encoding/csv"
	"errors"
	"stockels/graph/module/stock"
	"stockels/models"
	"stockels/utils"
	"strconv"

	"github.com/google/uuid"
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

func SubscribeMultipleStockService(subscribtions []models.Subscribtion, user models.User) ([]models.Subscribtion, error) {
	subStock := []models.Subscribtion{}

	for _, sub := range subscribtions {
		subscribtion := models.Subscribtion{
			StockSymbol: sub.StockSymbol,
			UserID: user.ID,
			SupportPrice: sub.SupportPrice,
			ResistancePrice: sub.ResistancePrice,
		}
		// err := utils.DB().Create(&subscribtion).Error
		err := utils.DB().Where(models.Subscribtion{StockSymbol: sub.StockSymbol, UserID: user.ID}).Assign(subscribtion).FirstOrCreate(&subscribtion).Error
		if err != nil {
			return subStock, err
		}
		subStock = append(subStock, subscribtion)

	}

	if len(subStock) == 0 {
		return subStock, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	return subStock, nil
}

func GetSubscribtionStockService(user models.User) ([]SubscribtionStockType, error) {
	subscribtions := []models.Subscribtion{}

	err :=  utils.DB().Find(&subscribtions, "user_id = ?", user.ID).Error
	if err != nil {
		return []SubscribtionStockType{}, err
	}

	subStock := []SubscribtionStockType{}

	for _, sub := range subscribtions {

		stock, err := stock.GetStockBySymbolService(sub.StockSymbol)
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
		return []SubscribtionStockType{}, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	return subStock, nil
}

func GenerateStockReportService(user models.User) (string, error) {
	subscribtions := []models.Subscribtion{}

	err :=  utils.DB().Find(&subscribtions, "user_id = ?", user.ID).Error
	if err != nil {
		return "", err
	}

	subStock := []SubscribtionStockType{}

	for _, sub := range subscribtions {

		stock, err := stock.GetStockBySymbolService(sub.StockSymbol)
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
		return "", errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	stocksRecords := [][]string{
		{"symbol", "name", "sector", "supportPercentage", "resistancePercentage", "supportPrice", "resistancePrice", "openPrice", "closePrice", "highestPrice", "lowestPrice", "volume", "lastUpdate", "website", "description"},
	}

	for _, record := range subStock {
		stocksRecords = append(stocksRecords, []string{record.Symbol, record.Name, record.Sector, PercentageFormat(record.SupportPercentage), PercentageFormat(record.ResistancePercentage), strconv.Itoa(record.SupportPrice), strconv.Itoa(record.ResistancePrice), record.OpenPrice, record.ClosePrice, record.HighestPrice, record.LowestPrice, record.Volume, record.LastUpdate, record.Website, record.Description})
	}

	stocksRecords = append(stocksRecords, )

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(stocksRecords) 

	fileName := uuid.New().String() + ".csv"
	reportFile, err := utils.FileUploader(csvBuffer, fileName);

	fileUrl := "https://stockels.s3.ap-southeast-1.amazonaws.com/" + *reportFile.Key
	return fileUrl, err
}

func GetReportStockService(user models.User, stocksReq []models.Subscribtion) (*bytes.Buffer, error) {
	subStock := []SubscribtionStockType{}

	for _, sub := range stocksReq {

		stock, err := stock.GetStockBySymbolService(sub.StockSymbol)
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

	stocksRecords = append(stocksRecords, )

	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(stocksRecords) 

	return csvBuffer, nil
}

func PercentageFormat(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 64)
}