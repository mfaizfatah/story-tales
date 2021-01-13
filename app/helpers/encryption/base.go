package encryption

import (
	"encoding/hex"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/scrypt"
)

// CipherKey key must be 32 chars long because block size is 16 bytes
var CipherKey = os.Getenv("SALT_KEY")

// AlgorithmNonceSize length of key
const AlgorithmNonceSize int = 12

const customchars = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-+`

var RdS = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func HashKeys(k string) string {
	salt := []byte(os.Getenv(`SALT_D`))
	keys := []byte(k)

	h, _ := scrypt.Key(keys, salt, 16384, 8, 1, 32)
	return hex.EncodeToString(h)
}

func SaltX() string {
	salt := []byte(os.Getenv(`SALT_C`))
	keys := []byte(os.Getenv(`SALT_A`))

	h, _ := scrypt.Key(keys, salt, 1024, 8, 1, 12)
	return hex.EncodeToString(h)
}

func RdAlpnum(length int) string {

	b := make([]byte, length)

	for i := range b {
		b[i] = customchars[RdS.Intn(len(customchars))]
	}
	return string(b)
}
