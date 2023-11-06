package analytic

import (
	"context"
	"encoding/json"
	"stockels/app/object"
	"stockels/app/utils"
	"strconv"
	"time"
)

func GetQuarterSupportAndResistance(quarter []*object.StockDetailPrice, quarterSymbol string) (*object.QuarterAnalytic)  {
	var supportPrice float64 = 9999999999
	var supportDate string
	var supportVolume int
	var resistancePrice float64
	var resistanceDate string
	var resistanceVolume int

	for _, eachPrice := range quarter {
		if ParseSrtingToFloat(eachPrice.Close) < supportPrice {
			supportPrice = ParseSrtingToFloat(eachPrice.Close)
			supportDate = eachPrice.Date
			supportVolume = eachPrice.Volume
		}

		if ParseSrtingToFloat(eachPrice.Close) > resistancePrice {
			resistancePrice = ParseSrtingToFloat(eachPrice.Close)
			resistanceDate = eachPrice.Date
			resistanceVolume = eachPrice.Volume
		}
	}
	// return &supportPrice, &supportDate, &supportVolume, &resistancePrice, &resistanceDate, &resistanceVolume

	return &object.QuarterAnalytic{Quarter: quarterSymbol, SupportPrice: supportPrice, SupportDate: supportDate, SupportVolume: supportVolume, ResistancePrice: resistancePrice, ResistanceDate: resistanceDate, ResistanceVolume: resistanceVolume,}
}

func CacheStockAnalytic(key string, stockAnalytic *object.Analytic) error {
	ctx := context.Background()
	stockAnalyticStringified, err := json.Marshal(stockAnalytic)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockAnalyticStringified, time.Hour).Err()

	return err
}

func ParseSrtingToFloat(payload string) (float64) {
	supportPrice, err := strconv.ParseFloat(payload, 64);

	if err != nil {
		panic(err.Error())
	}
	return supportPrice
}