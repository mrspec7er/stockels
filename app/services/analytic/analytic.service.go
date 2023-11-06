package analytic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"stockels/app/object"
	"stockels/app/services/stock"
	"stockels/app/utils"
	"strconv"
	"sync"
	"time"
)

func GetAnalyticFromAPI(symbol string) (*object.Analytic, error){
	analyticData := &object.Analytic{}

	cachedKey := symbol + "_Analytic"
	ctx := context.Background()
	cachedStockAnalytic, err := utils.Cache().Get(ctx, cachedKey).Result()
	if err != nil {	
		fmt.Println("Fetching business news to serpapi.com")
		upstreamSerpApiKey := os.Getenv("UPSTREAM_SERP_API_KEY")
	
		res, err := http.Get("https://serpapi.com/search.json?engine=google_finance&q="+ symbol+ ":IDX&api_key="+ upstreamSerpApiKey)
		if err != nil {
			return analyticData, err
		}
	
		analyticStreamMetaData, err := io.ReadAll(res.Body)
		if err != nil {
			return analyticData, err
		}
		
		err = json.Unmarshal(analyticStreamMetaData, &analyticData)
		if err != nil {
			return analyticData, err
		}
		err = CacheStockAnalytic(cachedKey, analyticData)
		return analyticData, err
	}

	err = json.Unmarshal([]byte(cachedStockAnalytic), &analyticData)
	
	return analyticData, err
}

func GetAnalyticStock(symbol string, fromYear int) (*object.TechnicalAnalytic, error){
	quarters := []*object.QuarterAnalytic{}
	toYear := time.Now().Year()

	quartersCtx := make(chan []*object.QuarterAnalytic, toYear - fromYear + 1)
	wg := &sync.WaitGroup{}
	wg.Add(toYear - fromYear + 1)
	
	for i := fromYear; i <= toYear; i++ {
		go GetQuarterStockPrice(symbol, i, quartersCtx, wg)
	}

	wg.Wait()
	close(quartersCtx)

	for eactQuarter := range quartersCtx {
		quarters = append(quarters, eactQuarter...)
	}

	if len(quarters) == 0 {
		return nil, errors.New("Failed to get data from 'GetQuarterStockPrice'!")
	}

	totalSupportPrice := 0
	totalResistancePrice := 0

	for _, qtr := range quarters {
		totalSupportPrice += int(qtr.SupportPrice)
		totalResistancePrice += int(qtr.ResistancePrice)
	}

	sort.Slice(quarters, func(i, j int) bool {
		return quarters[i].Quarter < quarters[j].Quarter
	  })

	return &object.TechnicalAnalytic{Quarters: quarters, AverageSupportPrice: float64(totalSupportPrice / len(quarters)), AverageResistancePrice: float64(totalResistancePrice/ len(quarters))}, nil
}

func GetQuarterStockPrice(symbol string, year int, quartersCtx chan []*object.QuarterAnalytic, wg *sync.WaitGroup)  {
	quarters := []*object.QuarterAnalytic{}
	
		firstQuarter := []*object.StockDetailPrice{}
		secondQuarter := []*object.StockDetailPrice{}
		thirdQuarter := []*object.StockDetailPrice{}
		fourthQuarter := []*object.StockDetailPrice{}

		stockPrices, err := stock.GetStockPriceFromAPI(symbol, strconv.Itoa(year) + "-01-01", strconv.Itoa(year) + "-12-31")
		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, priceDetail := range stockPrices {
			priceMounth := priceDetail.Date[5:7]
			if (priceMounth == "01" || priceMounth == "02" || priceMounth == "03") {
				firstQuarter = append(firstQuarter, priceDetail)
			}
			if (priceMounth == "04" || priceMounth == "05" || priceMounth == "06") {
				secondQuarter = append(secondQuarter, priceDetail)
			}
			if (priceMounth == "07" || priceMounth == "08" || priceMounth == "09") {
				thirdQuarter = append(thirdQuarter, priceDetail)
			}
			if (priceMounth == "10" || priceMounth == "11" || priceMounth == "12") {
				fourthQuarter = append(fourthQuarter, priceDetail)
			}
		}

		firstQuarterAnalytic := GetQuarterSupportAndResistance(firstQuarter, strconv.Itoa(year) + "-Q1")
		secondQuarterAnalytic:= GetQuarterSupportAndResistance(secondQuarter, strconv.Itoa(year) + "-Q2")
		thirdQuarterAnalytic:= GetQuarterSupportAndResistance(thirdQuarter, strconv.Itoa(year) + "-Q3")
		fourthQuarterAnalytic  := GetQuarterSupportAndResistance(fourthQuarter, strconv.Itoa(year) + "-Q4")

		quarters = append(quarters, firstQuarterAnalytic)
		quarters = append(quarters, secondQuarterAnalytic)
		quarters = append(quarters, thirdQuarterAnalytic)
		quarters = append(quarters, fourthQuarterAnalytic)
	
	quartersCtx <- quarters
	wg.Done()
	return
}

