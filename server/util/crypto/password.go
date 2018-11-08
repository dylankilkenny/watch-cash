package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func Encrypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func Compare(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
