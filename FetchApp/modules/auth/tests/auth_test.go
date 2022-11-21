package tests

import (
	"fetchapp/modules/auth"
	"fetchapp/modules/auth/models"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestAuth(t *testing.T) {
	x := rand.Intn(9)
	y := rand.Intn(9)
	z := rand.Intn(9)

	var phone, name, role string = fmt.Sprintf("+6287188723%d%d%d", x, y, z), "Jaya", "admin"

	os.Setenv("authapp_url", "http://localhost:8040")

	pass, err := auth.Register(phone, name, role)
	if err != nil {
		t.Fatalf("failed to register %s", err.Error())
	}

	fmt.Println(pass)

	tokenstring, err := auth.SignIn(phone, pass)
	if err != nil {
		t.Fatalf("failed to sign in %s", err.Error())
	}

	valid, err := auth.TokenClaim(tokenstring)
	if err != nil {
		t.Fatalf("failed to claim token %s", err.Error())
	} else if err == nil && (valid == models.UserDTO{}) {
		t.Fatalf("failed to validate, return empty")
	}
}
