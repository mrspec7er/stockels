package analytic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockels/app/object"
	"stockels/app/utils"
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

func CacheStockAnalytic(key string, stockAnalytic *object.Analytic) error {
	ctx := context.Background()
	stockAnalyticStringified, err := json.Marshal(stockAnalytic)

	if err != nil {
		return err
	}
	err = utils.Cache().Set(ctx, key, stockAnalyticStringified, time.Hour).Err()

	return err
}