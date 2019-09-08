package sha

import "testing"

func TestSHA1Encrypt(t *testing.T) {
	plain := "123456"
	t.Logf("Plain text : %s", plain)
	enc := SHA1Encrypt(plain)
	t.Logf("Encrypted : %s", enc)
}

func TestSHA256Encrypt(t *testing.T) {
	plain := "123456"
	t.Logf("Plain text : %s", plain)
	enc := SHA256Encrypt(plain)
	t.Logf("Encrypted : %s", enc)
}
