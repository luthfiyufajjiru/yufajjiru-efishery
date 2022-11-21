package models

type CurrencyExchangeDTO struct {
	Origin    int `json:"origin,omitempty"`
	Reference int `json:"reference,omitempty"`
}
