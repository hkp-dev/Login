package login

import (
	"crypto/sha256"
	"encoding/hex"
)

func AuthenticateUser(username, password string) bool {
	hashedPassword := HashPassword(password)

	for _, user := range Users {
		if user.Username == username && user.PasswordHash == hashedPassword {
			return true
		}
	}
	return false
}
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
