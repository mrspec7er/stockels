package subscribtion

import (
	"bytes"
	"encoding/csv"
	"errors"
	"stockels/app/models"
	"stockels/app/object"
	"stockels/app/services/stock"
	"stockels/app/utils"
	"strconv"

	"github.com/google/uuid"
)

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