package subscribtion

import (
	"bytes"
	"encoding/csv"
	"errors"
	"stockels/graph/module/stock"
	"stockels/graph/object"
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

func SubscribeMultipleStockService( stocks []*object.GetStockData, user *models.User) ([]*object.Subscribtion, error) {
	subs := []*object.Subscribtion{}

	for _, sub := range stocks {
		subscribtion := models.Subscribtion{
			StockSymbol: sub.StockSymbol,
			UserID: user.ID,
			SupportPrice: sub.SupportPrice,
			ResistancePrice: sub.ResistancePrice,
		}
		// err := utils.DB().Create(&subscribtion).Error
		err := utils.DB().Where(models.Subscribtion{StockSymbol: sub.StockSymbol, UserID: user.ID}).Assign(subscribtion).FirstOrCreate(&subscribtion).Error
		if err != nil {
			return subs, err
		}
		subs = append(subs, &object.Subscribtion{StockSymbol: subscribtion.StockSymbol, UserID: user.ID, SupportPrice: subscribtion.SupportPrice, ResistancePrice: subscribtion.ResistancePrice})

	}

	if len(subs) == 0 {
		return subs, errors.New("Failed to get data from 'GetStockBySymbolService'!")
	}

	return subs, nil
}

func GetSubscribtionStockService(user models.User) ([]*object.StockData, error) {
	subscribtions := []models.Subscribtion{}
	stocks := []*object.StockData{}
	stocksReq := []*object.GetStockData{}

	err :=  utils.DB().Find(&subscribtions, "user_id = ?", user.ID).Error
	if err != nil {
		return stocks, err
	}


	for _, subs := range subscribtions {
		stocksReq = append(stocksReq, &object.GetStockData{StockSymbol: subs.StockSymbol, SupportPrice: subs.SupportPrice, ResistancePrice: subs.ResistancePrice})
	}

	stocks, err = stock.GetMultipleStockService(stocksReq)

	if err != nil {
		return stocks, errors.New("Failed to get data from 'GetMultipleStockService'!")
	}

	return stocks, nil
}

func GenerateStockReportService(user models.User) (*object.GenerateReportResponse, error) {
	subscribtions := []models.Subscribtion{}
	response := &object.GenerateReportResponse{}
	subStock := []*object.StockData{}
	stocksReq := []*object.GetStockData{}

	err :=  utils.DB().Find(&subscribtions, "user_id = ?", user.ID).Error
	if err != nil {
		return response, err
	}


	for _, subs := range subscribtions {
		stocksReq = append(stocksReq, &object.GetStockData{StockSymbol: subs.StockSymbol, SupportPrice: subs.SupportPrice, ResistancePrice: subs.ResistancePrice})
	}

	subStock, err = stock.GetMultipleStockService(stocksReq)

	if err != nil {
		return response, errors.New("Failed to get data from 'GetMultipleStockService'!")
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

	fileName := uuid.New().String() + ".csv"
	reportFile, err := utils.FileUploader(csvBuffer, fileName);

	fileUrl := "https://stockels.s3.ap-southeast-1.amazonaws.com/" + *reportFile.Key
	return &object.GenerateReportResponse{ReportURL: fileUrl}, err
}

func PercentageFormat(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', 2, 64)
}