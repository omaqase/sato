package payment

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var encryptionKey = []byte("your-32-byte-encryption-key-here") // В реальном проекте должен быть в конфиге

func EncryptCardNumber(cardNumber string) []byte {
	return Encrypt([]byte(cardNumber))
}

func EncryptCVV(cvv string) []byte {
	return Encrypt([]byte(cvv))
}

func Encrypt(data []byte) []byte {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {

		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	return gcm.Seal(nonce, nonce, data, nil)
}
