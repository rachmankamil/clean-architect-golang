package encrypt

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	return string(bytes), err
}

func ValidateHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
