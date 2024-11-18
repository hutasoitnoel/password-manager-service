package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"password-manager-service/config"
)

// Decrypt decrypts an AES-encrypted string or returns the plain text if it wasn't encrypted.
func Decrypt(encryptedText string) (string, error) {
	fmt.Printf("Encrypted text length: %d\n", len(encryptedText))

	// Try decoding Base64
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		fmt.Println("Not a valid Base64 string, returning as plain text")
		return encryptedText, nil // Treat it as plain text
	}

	fmt.Printf("Decoded cipherText length: %d\n", len(cipherText))
	fmt.Println(config.ENCRYPTION_KEY)

	// Create AES block cipher
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
		fmt.Println("CipherText too short, returning as plain text")
		return encryptedText, nil // Treat it as plain text
	}

	// Split nonce and actual cipher text
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	// Attempt decryption
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Printf("Decryption failed: %v, returning as plain text\n", err)
		return encryptedText, nil // Treat it as plain text
	}

	return string(plainText), nil
}
