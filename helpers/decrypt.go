package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"password-manager-service/config"
)

// Decrypt decrypts an AES-encrypted string.
func Decrypt(encryptedText string) (string, error) {
	fmt.Printf("Encrypted text length: %d\n", len(encryptedText))

	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	fmt.Printf("Decoded cipherText length: %d\n", len(cipherText))

	fmt.Println(config.ENCRYPTION_KEY)

	block, err := aes.NewCipher(config.ENCRYPTION_KEY)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
