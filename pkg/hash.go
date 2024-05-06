package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hasedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hasedBytes), nil
}

func VerifyPassword(hashPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
}
