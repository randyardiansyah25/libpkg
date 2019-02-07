package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"github.com/randyardiansyah25/libpkg/util/str"
)

func Encrypt(key, iv, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := text

	b = PKCS5Padding(b, aes.BlockSize, len(text))
	ciphertextb := make([]byte, len(b))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertextb, b)
	return strutils.ByteToHexString(ciphertextb), nil
}

func Decrypt(key []byte, iv []byte, encText string) (string, error) {
	text, err := strutils.HexStringToByte(encText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, text)

	return string(PKCS5UnPadding(decrypted)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
