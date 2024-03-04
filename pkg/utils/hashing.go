package utils

import "golang.org/x/crypto/bcrypt"

func HashValue(v string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CompareToHash(v, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(v))
	return err == nil
}
