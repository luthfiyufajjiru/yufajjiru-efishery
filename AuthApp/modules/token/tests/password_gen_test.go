package tests

import (
	"authapp/modules/token/services"
	"testing"
)

func TestPasswordGen(t *testing.T) {
	x := services.DefaultPassword()

	for i := 0; i < 10; i++ {
		xs := services.DefaultPassword()
		if x == xs {
			t.Log("not unique")
			t.Fail()
		}

		if len(xs) != 4 {
			t.Log("not expected")
			t.Fail()
		}

		x = xs
	}
}
