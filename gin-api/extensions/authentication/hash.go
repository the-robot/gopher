package authentication

import (
	"golang.org/x/crypto/bcrypt"

	"gingo/extensions/error"
)

func HashPassword(password string) (string, error.IError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), error.Internal(err.Error(), err)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
