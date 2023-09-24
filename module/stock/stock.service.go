package stock

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"stockels/models"
	"stockels/utils"
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

type GoapiResponseType struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data struct {
		Result *StockDataType `json:"result"`
		LastPrice *StockPriceType `json:"last_price"`
	} `json:"data"`
}

func GetAllStockServices(subscribtions []models.Subscribtion) ([]models.Stock, error) {
	stocks := []models.Stock{}
	for _, sub := range subscribtions {

		stock, err := GetEachStockServices(sub.StockSymbol)
		if err != nil {
			break
		}
		stocks = append(stocks, stock)
	}

	if len(stocks) == 0 {
		return stocks, errors.New("Invalid stock symbol!")
	}

	return stocks, nil
}

func GetEachStockServices(symbol string) (models.Stock, error) {
	stockMetaData := GoapiResponseType{}
	res, err := http.Get("https://api.goapi.id/v1/stock/idx/" + symbol + "?api_key=uz0801JrrjNL0sAGpDCSvNzAvj2lBL")
	if err != nil {
		return models.Stock{}, err
	}
	stockStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Stock{}, err
	}
	err = json.Unmarshal(stockStreamMetaData, &stockMetaData)
	if err != nil || stockMetaData.Data.Result == nil {
		return models.Stock{}, errors.New("Invalid Symbol")
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