package models

import (
	"fetchapp/modules/currency"
	"fmt"
	"strconv"
)

// uuid: "b175e2c0-a97e-4cf4-b16f-ab8f59c1a429",
// komoditas: "GURAME",
// area_provinsi: "KALIMANTAN TIMUR",
// area_kota: "BORNEO",
// size: "30",
// price: "87000",
// tgl_parsed: "2022-01-01T23:34:16Z",
// timestamp: "1641080056368"

type ResourceDTO struct {
	Uuid         string `json:"uuid,omitempty"`
	Komoditas    string `json:"komoditas,omitempty"`
	AreaProvinsi string `json:"area_provinsi,omitempty"`
	AreaKota     string `json:"area_kota,omitempty"`
	Size         string `json:"size,omitempty"`
	Price        string `json:"price,omitempty"`
	PriceUsd     string `json:"price_usd,omitempty"`
	TglParsed    string `json:"tgl_parsed,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}

func PriceUSD(price string) string {
	curr, _ := currency.GetExchange()

	x, _ := strconv.ParseFloat(price, 64)

	return fmt.Sprintf("%f", x/float64(curr.Origin))

}

func SetAllPriceUSD(s []ResourceDTO) []ResourceDTO {

	for i := range s {
		s[i].PriceUsd = PriceUSD(s[i].Price)
	}

	return s
}
