package tests

import (
	"authapp/modules/account/services"
	"testing"
)

func TestPhoneNumberValidation(t *testing.T) {
	phone_numbers := []string{"+6281344444444", "zzzzzz", "081344444444", "082212312"}
	res1 := services.PhoneNumberValidation(phone_numbers[0])
	if !res1 {
		t.Log("harus diawali +62")
		t.Fail()
	}

	res2 := services.PhoneNumberValidation(phone_numbers[1])
	if res2 {
		t.Log("harus diawali +62")
		t.Fail()
	}

	res3 := services.PhoneNumberValidation(phone_numbers[2])
	if res3 {
		t.Log("harus diawali +62")
		t.Fail()
	}

	res4 := services.PhoneNumberValidation(phone_numbers[3])
	if res4 {
		t.Log("harus diawali +62")
		t.Fail()
	}
}
