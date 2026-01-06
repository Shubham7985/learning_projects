package utils

import "golang.org/x/crypto/bcrypt"

/*
HashPassword
Plain password ko secure hash me badal deta hai
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
CheckPassword
Login ke time password match karta hai
*/
func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
