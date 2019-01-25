package aes

import (
	"testing"
)

func TestAesCrypto_Encrypt(t *testing.T) {
	key := []byte("us5i3channel!@@!")
	iv := key
	plaintext := "123456"

	t.Logf("Plain text : %s\n", plaintext)
	ciphertext, err := Encrypt(key, iv, []byte(plaintext))
	if err != nil {
		t.Error(err)
	}
	//t.Logf("Encrypted  : %0x\n", ciphertext)
	t.Logf("Encrypted  : %s\n", ciphertext)

}

func TestAesCrypto_Decrypt(t *testing.T) {
	key := []byte("us5i3channel!@@!")
	iv := key
	ciphertext := "2397140C989F8BBB061250F419E84D34"

	t.Logf("Encrypted text : %s\n", ciphertext)
	plaintextb, err := Decrypt(key, iv, ciphertext)
	if err != nil {
		t.Error(err)
	}
	plaintext := string(plaintextb)
	t.Logf("Encrypted  : %s\n", plaintext)
}
