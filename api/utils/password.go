package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword receives a password string and hashes it using using bcrypt
// Returns the hashedPassword as string and an error if the hash failed
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPasswordHash checks if a password is equal to the given hash
// Returns true if equal and false if not
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
