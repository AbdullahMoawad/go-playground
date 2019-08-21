// passwords.go
package common

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//func main() {
//	password := "secret"
//	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
//
//	fmt.Println("Password:", password)
//	fmt.Println("Hash:    ", hash)
//
//	match := CheckPasswordHash(password, hash)
//	fmt.Println("Match:   ", match)
//}