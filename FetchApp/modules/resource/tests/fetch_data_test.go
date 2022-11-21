package tests

import (
	"fetchapp/modules/resource/services"
	"testing"
)

func TestFetchResources(t *testing.T) {
	for i := 0; i < 2; i++ {
		d := services.GetResources()

		if d == nil {
			t.Fatal("failed to fetch resources")
		}
	}
}

func TestFetchResourcesUSD(t *testing.T) {
	d := services.GetResources()

	if d == nil {
		t.Fatal("failed to fetch resources")
	}

	c := services.SetAllPriceUSD(d)

	for i, v := range c {
		if v.Price != "" && v.PriceUsd == "" {
			t.Fatal("failed to convert")
		}

		if i > 3 {
			break
		}
	}

}
