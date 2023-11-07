package testings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"stockels/app/object"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStockController(t *testing.T)  {
	router := SetupRouters("../.env")
	w := httptest.NewRecorder()

	t.Run("should generate report files", func(t *testing.T) {

		body := []object.GetStockData{{StockSymbol: "ICBP", SupportPrice: 9300, ResistancePrice: 11500}}
	
		bodyJSON, err := json.Marshal(body)
		payload := bytes.NewBuffer(bodyJSON)
		if err != nil {
			   fmt.Printf("server: could not read request body: %s\n", err)
		}
	
		req, err := http.NewRequest("POST", "/api/v1/stocks/generate-report", payload)
		req.Header.Set("Content-Type", "application/json")
	
		if err != nil {
			t.Errorf(err.Error())
		}
	
		router.ServeHTTP(w, req)
	
		assert.Equal(t, 200, w.Code)
	
		if (reflect.TypeOf(w.Body) != reflect.TypeOf(bytes.NewBufferString(""))) {
			t.Errorf(err.Error())
		}
	})

}