package utils

import "golang.org/x/crypto/bcrypt"

// HashIsValid - Return true if the value matches the hashed value else return `false`
func HashIsValid(value string, hashedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))

	return err == nil
}

func MakeHash(password string) string {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPass)
}
