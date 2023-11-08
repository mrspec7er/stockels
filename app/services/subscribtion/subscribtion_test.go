package subscribtion_test

import (
	"reflect"
	"stockels/app/models"
	"stockels/app/object"
	"stockels/app/services/subscribtion"
	"stockels/testings"
	"testing"
)

func init()  {
	testings.SetupRouters("../../../.env")
}

func TestSubscribtionServices(t *testing.T)  {
	t.Run("should create new stock subscribtion", func(t *testing.T) {

		user := &models.User{ID: 1, FullName: "Mr Spec7er", Email: "mrspec7er@gmail.com", IsVerified: false}
		stocks := []*object.GetStockData{{StockSymbol: "BMRI", SupportPrice: 5450, ResistancePrice: 6700}}
		subscribtions := []*object.Subscribtion{}

		result, err := subscribtion.SubscribeMultipleStockService(stocks, user)
	
		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(subscribtions)) {
			t.Errorf(err.Error())
		}
	})
	
}