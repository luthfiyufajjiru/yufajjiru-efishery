package token

import (
	"fmt"
	"testing"
)

func TestJWTToken(t *testing.T) {

	phone, name, role := "+6281315809254", "Luthfi", "User"

	token_string, _ := GenerateJWT(phone, name, role)

	fmt.Println(token_string)

	claims, err := Claims(token_string)

	if err != nil || claims["name"] != name && claims["phone_number"] != phone {
		t.Log("claims not valid")
		t.Fail()
	}
}
