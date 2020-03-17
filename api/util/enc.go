package util

import (
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

func Enc(password, salt string) (ct string) {
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}
