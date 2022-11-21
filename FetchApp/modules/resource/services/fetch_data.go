package services

import (
	"encoding/json"
	"fetchapp/modules/caching"
	"fetchapp/modules/resource/models"
	"net/http"
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
