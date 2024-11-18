package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// Encrypt encrypts a plaintext string using AES.
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		return "", err
	}

	fmt.Println("creating new GCM")
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	fmt.Println("creating new NonceSize")
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	fmt.Println("creating new seal")
	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
