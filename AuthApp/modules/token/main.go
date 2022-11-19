package token

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secreet_key []byte

func init() {
	secreet_key = []byte(os.Getenv("ajhsidi"))
}

func GenerateJWT(phone_number, name, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["phone_number"] = phone_number
	claims["name"] = name
	claims["role"] = role
	claims["timestamp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(secreet_key)

	if err != nil {
		msg := fmt.Sprintf("Something Went Wrong: %s", err.Error())
		_ = fmt.Errorf(msg)
		return msg, err
	}
	return tokenString, nil
}

func Claims(token_input string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(
		token_input,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid input")
			}
			return secreet_key, nil
		})

	if token == nil {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("token is invalid")
	}

	return
}
