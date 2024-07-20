package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, cost)
	return string(bytes), err
}

func CheckHashPassword(hashPass []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashPass, password)
}
