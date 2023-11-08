package testings

import (
	"context"
	"reflect"
	"stockels/app"
	"stockels/app/handlers"
	"stockels/app/models"
	"stockels/app/object"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

// Mock request context
func addContext(user models.User) client.Option {
    return func(bd *client.Request) {
		// c := gin.Context{}
        ctx := context.WithValue(bd.HTTP.Context(), "user", user)
        bd.HTTP = bd.HTTP.WithContext(ctx)
    }
}

func TestSubscriptionResolvers(t *testing.T)  {
	SetupRouters("../.env")
	c := client.New(handler.NewDefaultServer(app.NewExecutableSchema(app.Config{Resolvers: &handlers.Resolver{}})))

	t.Run("should return array of subscribed stock", func(t *testing.T) {
		var stocks struct {
			GetStockSubscribe []*object.StockData
		}

		user := models.User{ID: 1, FullName: "Kusuma Sandi", Email: "mrspec7er@gmail.com", IsVerified: false}
		c.MustPost(`query {
			getStockSubscribe {
				supportPercentage,
				resistancePercentage,
				symbol,
				name,
				description,
				sector,
			  }
		}`, &stocks, addContext(user))
		
		if (reflect.TypeOf(stocks.GetStockSubscribe) != reflect.TypeOf([]*object.StockData{})) {
			t.Errorf("Invalid response, value are not equal to array of stockdata")
		}
	});

	t.Run("should return report url", func(t *testing.T) {
		var reportURL struct {
			GenerateReportFile *object.GenerateReportResponse
		}

		user := models.User{ID: 1, FullName: "Kusuma Sandi", Email: "mrspec7er@gmail.com", IsVerified: false}
		c.MustPost(`query {
			generateReportFile {reportUrl}
		}`, &reportURL, addContext(user))
		
		if (reflect.TypeOf(reportURL.GenerateReportFile.ReportURL).Kind() != reflect.String) {
			t.Errorf("Invalid response, url must be string")
		}
	});
}