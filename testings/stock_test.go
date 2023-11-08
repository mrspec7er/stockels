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
	var stockDetail struct {
		GetStockDetail *object.StockDetail
	}

	t.Run("should generate report files", func(t *testing.T) {
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

}