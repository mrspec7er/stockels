package testings

import (
	"stockels/app"
	"stockels/app/handlers"
	"stockels/app/object"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
)

func TestStockResolvers(t *testing.T)  {
	SetupRouters("../.env")
	c := client.New(handler.NewDefaultServer(app.NewExecutableSchema(app.Config{Resolvers: &handlers.Resolver{}})))
	
	t.Run("should return stock info", func(t *testing.T) {
		var stockDetail struct {
			GetStockDetail *object.StockDetail
		}
		c.MustPost(`query($supportPrice: Int!, $resistancePrice: Int!) {
			getStockDetail(
				symbol: "ASII",
				fromDate: "2023-10-05",
				toDate: "2023-10-11",
				supportPrice: $supportPrice,
				resistancePrice: $resistancePrice
			){
				info {symbol, name, supportPercentage, resistancePercentage, sector, closePrice, description},
				price {close, open, high, low, volume, date}
			}
		}`, &stockDetail, client.Var("supportPrice", 5300), client.Var("resistancePrice", 7200))
		
		assert.Equal(t, "ASII", stockDetail.GetStockDetail.Info.Symbol)
	});

	t.Run("should return array of stock info", func(t *testing.T) {
		var stocks struct {
			GetStocks []*object.StockData
		}
		c.MustPost(`query($supportPrice: Int!, $resistancePrice: Int!) {
			getStocks(
				stocks: [
						{stockSymbol: "ASII", supportPrice: $supportPrice, resistancePrice: $resistancePrice}, 
					]
			  ) {
				symbol
				name
				closePrice
				supportPercentage
				resistancePercentage,
					description
			  }
		}`, &stocks, client.Var("supportPrice", 5300), client.Var("resistancePrice", 7200))
		
		assert.Equal(t, "ASII", stocks.GetStocks[0].Symbol)
	});

}