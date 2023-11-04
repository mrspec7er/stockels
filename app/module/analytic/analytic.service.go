package analytic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockels/app/object"
)

func GetAnalyticFromAPI(symbol string) (*object.Analytic, error){
	fmt.Println("Fetching business news to serpapi.com")
	upstreamSerpApiKey := os.Getenv("UPSTREAM_SERP_API_KEY")
	analyticData := &object.Analytic{}

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
	
	return analyticData, err
}