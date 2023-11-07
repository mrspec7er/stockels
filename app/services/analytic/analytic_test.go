package analytic_test

import (
	"reflect"
	"stockels/app/object"
	"stockels/app/services/analytic"
	"stockels/testings"
	"testing"
)

func init()  {
	testings.SetupRouters("../../../.env")
}

func TestAnalyticServices(t *testing.T)  {
	t.Run("should return stock fundamental analytic", func(t *testing.T) {
		stockAnalytic := &object.Analytic{}

		result, err := analytic.GetAnalyticFromAPI("TLKM")
	
		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stockAnalytic)) {
			t.Errorf(err.Error())
		}
	})

	t.Run("should return stock technical analytic", func(t *testing.T) {
		stockTechnicalAnalytic := &object.TechnicalAnalytic{}

		result, err := analytic.GetAnalyticStock("ASII", 2021)
	
		if err != nil || (reflect.TypeOf(result) != reflect.TypeOf(stockTechnicalAnalytic)) {
			t.Errorf(err.Error())
		}
	})
	
}