package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failedto hash password")
	}
	return string(hashedPassword),
		nil
}

func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func PasswordEncoder(text string) string {
	md5 := Md5Hash(text)

	randomString := ""
	for i := 0; i < len(md5); i++ {
		if i%2 == 0 {
			randomString += string(md5[i])
		}
	}

	resultBase64 := base64.StdEncoding.EncodeToString([]byte(randomString))

	finalHash := Md5Hash(resultBase64)

	return finalHash
}
