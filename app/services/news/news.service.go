package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"stockels/app/object"
)

type NewsapiPriceResponseType struct {
	TotalResults int `json:"totalResults"`
	Articles []*object.Article `json:"articles"`
}

func GetNewsFromAPI() ([]*object.Article, error){
	fmt.Println("Fetching business news to newsapi.org")
	newsMetaData := NewsapiPriceResponseType{}
	upstreamNewsApiKey := os.Getenv("UPSTREAM_NEWS_API_KEY")

	res, err := http.Get("https://newsapi.org/v2/top-headlines?country=id&category=business&apiKey=" + upstreamNewsApiKey)
	if err != nil {
		return []*object.Article{}, err
	}

	newsStreamMetaData, err := io.ReadAll(res.Body)
	if err != nil {
		return []*object.Article{}, err
	}

	err = json.Unmarshal(newsStreamMetaData, &newsMetaData)
	if err != nil || newsMetaData.Articles == nil {
		return []*object.Article{}, errors.New("Failed to fetch price data from goapi.id")
	}

	news := newsMetaData.Articles
	
	return news, err
}