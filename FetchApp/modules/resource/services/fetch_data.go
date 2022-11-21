package services

import (
	"encoding/json"
	"fetchapp/modules/caching"
	"fetchapp/modules/currency"
	"fetchapp/modules/resource/models"
	"fmt"
	"net/http"
	"strconv"
)

var data caching.SessionCaching

func init() {
	data = caching.SessionCaching{
		FetchData: getResources,
	}

	data.Begin()
}

func getResources() interface{} {
	url := "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"

	var results []models.ResourceDTO

	client, err := http.Get(url)
	if x, y := json.NewDecoder(client.Body).Decode(&results), client.Body.Close(); x != nil || y != nil || err != nil {
		results = nil
	}

	return results
}

func GetResources() []models.ResourceDTO {
	x, _ := data.GetData()
	y := *x
	z := y.([]models.ResourceDTO)

	return z
}

func PriceUSD(price string) string {
	curr, _ := currency.GetExchange()

	x, _ := strconv.ParseFloat(price, 64)

	return fmt.Sprintf("%f", x/float64(curr.Origin))

}

func SetAllPriceUSD(s []models.ResourceDTO) []models.ResourceDTO {

	for i := range s {
		if s[i].Price != "" {
			s[i].PriceUsd = PriceUSD(s[i].Price)
		}
	}

	return s
}
