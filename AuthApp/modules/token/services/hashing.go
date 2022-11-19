package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltString(inp string) []byte {
	res, _ := bcrypt.GenerateFromPassword([]byte(inp), bcrypt.MinCost)
	return res
}

func CompareHashAndString(stored []byte, inp string) bool {
	err := bcrypt.CompareHashAndPassword(stored, []byte(inp))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
