package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	_ "github.com/joho/godotenv/autoload" //autoload env
)

// Encrypt encrypts plain text string into cipher text string
func Encrypt(plaintext []byte) (string, error) {
	key := []byte(CipherKey)
	// Generate a 96-bit nonce using a CSPRNG.
	nonce := make([]byte, AlgorithmNonceSize)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}

	// Create the cipher and block.
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt and prepend nonce.
	ciphertext := cipher.Seal(nil, nonce, plaintext, nil)
	ciphertextAndNonce := make([]byte, 0)

	ciphertextAndNonce = append(ciphertextAndNonce, nonce...)
	ciphertextAndNonce = append(ciphertextAndNonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(ciphertextAndNonce), nil
}

// EncryptorTokenForgotPass is a func for encrypt token forgot password
func EncryptorTokenForgotPass(text string) string {

	plain := []byte(text)
	keys, _ := hex.DecodeString(HashKeys(`SALT_B`))

	block, _ := aes.NewCipher(keys)
	nonce, _ := hex.DecodeString(SaltX())
	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, plain, nil)

	return hex.EncodeToString(ciphertext)
}
