package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(password string) ([]byte, error) {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Println("Error occoured while generating hash")
		return nil, err
	}
	return hash, nil
}

func DecryptPass(hash string) {

}
