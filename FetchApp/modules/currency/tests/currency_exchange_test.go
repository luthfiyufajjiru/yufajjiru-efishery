package tests

import (
	"fetchapp/modules/currency"
	"log"
	"testing"
)

func TestGetExchange(t *testing.T) {
	for i := 0; i < 3; i++ {
		x, err := currency.GetExchange()

		if err != nil {
			t.Fatal("failed to fetch")
		}

		log.Printf("Origin was %d and refference was %d", x.Origin, x.Reference)
	}

}
