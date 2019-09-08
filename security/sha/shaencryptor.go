package sha

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func SHA1Encrypt(plain string) string {
	h := sha1.New()
	h.Write([]byte(plain))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func SHA256Encrypt(plain string) string {
	hash := sha256.New()
	hash.Write([]byte(plain))
	bs := hash.Sum(nil)
	return hex.EncodeToString(bs)
}
