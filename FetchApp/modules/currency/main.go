package currency

import (
	"fetchapp/modules/caching"
	"fetchapp/modules/currency/models"
	"time"
)

var (
	data caching.SessionCaching
)

func getCurrencyExchange() interface{} {
	time.Sleep(time.Millisecond * 1300)
	return models.CurrencyExchangeDTO{
		Origin:    15000,
		Reference: 1,
	}
}

func init() {
	data = caching.SessionCaching{
		FetchData: getCurrencyExchange,
	}
	data.Begin()
}

func GetExchange() (*models.CurrencyExchangeDTO, error) {
	x, err := data.GetData()
	y := *x
	z := y.(models.CurrencyExchangeDTO)
	return &z, err
}
