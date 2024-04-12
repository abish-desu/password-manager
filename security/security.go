package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func EncryptPassword(password []byte, cipherKey string) []byte {
	aesBlock, err := aes.NewCipher([]byte(createHash(cipherKey)))
	if err != nil {
		fmt.Println(err)

	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	cipherText := gcmInstance.Seal(nonce, nonce, password, nil)

	return cipherText
}
func DecryptPassword(hashedPassword []byte, cipherKey string) []byte {
	aesBlock, err := aes.NewCipher([]byte(createHash(cipherKey)))
	if err != nil {
		fmt.Println(err)

	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := hashedPassword[:nonceSize], hashedPassword[nonceSize:]
	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		fmt.Println(err)
	}
	return originalText
}
