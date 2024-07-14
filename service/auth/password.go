package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string{
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePassword(hashed string, plain[] byte) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}