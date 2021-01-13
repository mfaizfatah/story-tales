package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"

	_ "github.com/joho/godotenv/autoload" //autoload env
)

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string) ([]byte, error) {
	key := []byte(CipherKey)
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	// Create slices pointing to the ciphertext and nonce.
	nonce := cipherText[:AlgorithmNonceSize]
	ciphertext := cipherText[AlgorithmNonceSize:]

	// Create the cipher and block.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt and return result.
	plaintext, err := cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// DecryptorTokenForgotPass is a func for decrypt token forgot password
func DecryptorTokenForgotPass(text string) (string, error) {

	plain, _ := hex.DecodeString(text)
	keys, _ := hex.DecodeString(HashKeys(`SALT_B`))
	block, _ := aes.NewCipher(keys)
	aesgcm, _ := cipher.NewGCM(block)
	nonce, _ := hex.DecodeString(SaltX())

	plaintext, err := aesgcm.Open(nil, nonce, plain, nil)
	if err != nil {
		return string(plaintext), err
	}

	return string(plaintext), nil
}
