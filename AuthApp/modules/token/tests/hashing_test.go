package tests

import (
	"authapp/modules/token/services"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println("Test hashing and validate hash")

	msg := "Hello world"
	hashed := services.HashAndSaltString(msg)

	test1 := services.CompareHashAndString(hashed, msg)
	test2 := services.CompareHashAndString(hashed, "hello")

	if !test1 {
		t.Logf("should got: true but got: %t", test1)
		t.Fail()
	}

	if test2 {
		t.Logf("should got: false but got: %t", test2)
		t.Fail()
	}

	t.Log("Success")
}
