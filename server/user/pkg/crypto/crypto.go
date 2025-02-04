package crypto

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func GenerateHashFromPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("error happened while hashing with bcrypt: %s", err.Error())
	}

	return string(hash), nil
}

func GenerateOTPCode() (string, error) {
	min, max := int64(100000), int64(999999)
	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP code: %w", err)
	}
	return fmt.Sprintf("%06d", min+n.Int64()), nil
}
